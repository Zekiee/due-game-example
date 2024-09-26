import ByteBuffer from 'bytebuffer';
import { Packer, Message } from "./packer";
// const WebSocket = require('ws');

export interface ClientOptions {
    // 连接地址
    url: string;
    // 打包器
    packer: Packer;
    // 心跳间隔，大于0则开启心跳
    heartbeat: number;
}

export interface ConnectHandler {
    (): any
}

export interface DisconnectHandler {
    (): any
}

export interface ReceiveHandler {
    (message: Message): any
}

export interface ErrorHandler {
    (): any
}

export interface HeartbeatHandler {
    (millisecond?: number): any
}

export class Client {
    // 连接打开hook函数
    private connectHandler?: ConnectHandler;
    // 连接关闭hook函数
    private disconnectHandler?: DisconnectHandler;
    // 接收消息hook函数
    private receiveHandler?: ReceiveHandler;
    // 错误处理hook函数
    private errorHandler?: ErrorHandler;
    // 心跳处理hook函数
    private heartbeatHandler?: HeartbeatHandler;
    // 客户端配置
    private opts: ClientOptions;
    // websocket
    private websocket?: WebSocket;
    // 定时器ID
    private intervalId: any;
    // 打包器
    private packer: Packer;
    // 
    private buffer: any;
    // 同步等待组
    private waitgroup: Map<number, {
        seq: number;
        callback: Map<number, (message: Message) => void>
    }>;

    public constructor(opts: ClientOptions) {
        this.opts = opts;
        this.websocket = undefined;
        this.packer = opts.packer || new Packer(opts.packer);
        this.waitgroup = new Map();
        this.buffer = new ByteBuffer(ByteBuffer.DEFAULT_CAPACITY, false, ByteBuffer.DEFAULT_NOASSERT);
    }

    /**
     * 连接服务区
     * @returns boolean
     */
    public connect(): boolean {
        try {
            this.disconnect();

            // 新建WS客户端
            this.websocket = new WebSocket(this.opts.url);
            this.websocket.binaryType = 'arraybuffer';

            // 监听WS连接打开
            this.websocket.onopen = (ev: Event) => {
                this.heartbeat();

                this.connectHandler && this.connectHandler();
            }

            // 监听WS连接关闭
            this.websocket.onclose = (ev: CloseEvent) => {
                this.disconnectHandler && this.disconnectHandler();
            }

            // 监听WS连接错误
            this.websocket.onerror = (ev: Event) => {
                this.errorHandler && this.errorHandler();
            }

            // 监听WS消息
            this.websocket.onmessage = (e: MessageEvent) => {
                if (e.data.byteLength == 0) {
                    return;
                }

                const packet = this.packer.unpack(e.data);

                if (packet.isHeartbeat) {
                    this.heartbeatHandler && this.heartbeatHandler(packet.millisecond);
                } else {
                    packet.message && (this.invoke(packet.message) || this.receiveHandler && this.receiveHandler(packet.message));
                }
            }

            return true
        } catch (error) {
            console.log(error)
            return false
        }
    }

    /**
     * 断开连接
     * @returns void
     */
    public disconnect() {
        if (!this.websocket) {
            return
        }

        this.intervalId && clearInterval(this.intervalId);
        this.intervalId = null;
        const websocket = this.websocket;
        this.websocket = undefined;
        const onclose = websocket.onclose;
        const onempty = () => { };
        websocket.onopen = onempty;
        websocket.onmessage = onempty;
        websocket.onclose = onempty;
        websocket.onerror = onempty;
        websocket.close();
    }

    /**
     * 启动心跳
     * @returns void
     */
    private heartbeat() {
        if (!this.opts.heartbeat || this.opts.heartbeat <= 0) {
            return;
        }

        const data = this.packer.packHeartbeat();

        this.intervalId = setInterval(() => {
            this.isConnected() && this.websocket && this.websocket.send(data);
        }, this.opts.heartbeat);
    }

    /**
     * 设置连接处理器
     * @param handler ConnectHandler
     */
    public onConnect(handler: ConnectHandler) {
        this.connectHandler = handler
    }

    /**
     * 设置断开连接处理器
     * @param handler DisconnectHandler
     */
    public onDisconnect(handler: DisconnectHandler) {
        this.disconnectHandler = handler
    }

    /**
     * 设置消息接收处理器
     * @param handler ReceiveHandler
     */
    public onReceive(handler: ReceiveHandler) {
        this.receiveHandler = handler
    }

    /**
     * 设置错误处理器
     * @param handler ErrorHandler
     */
    public onError(handler: ErrorHandler) {
        this.errorHandler = handler
    }

    /**
     * 设置心跳处理器
     * @param handler HeartbeatHandler
     */
    public onHeartbeat(handler: HeartbeatHandler) {
        this.heartbeatHandler = handler
    }

    /**
     * 检测客户端是否已连接
     * @returns boolean
     */
    public isConnected(): boolean {
        return this.websocket !== undefined && this.websocket.readyState == WebSocket.OPEN
    }

    /**
     * 检测客户端是否正在连接
     * @returns boolean
     */
    public isConnecting(): boolean {
        return this.websocket !== undefined && this.websocket.readyState === WebSocket.CONNECTING
    }

    /**
     * 发送消息
     * @param message Message 消息
     * @returns boolean
     */
    public send(message: Message): boolean {
        if (this.isConnected()) {
            const data = this.packer.packMessage(message);
            
            this.websocket && this.websocket.send(data);

            return true;
        }

        return false;
    }

    /**
     * 请求C/S模型，需要客户端和服务器同时开启seq配置支持
     * @param route 路由
     * @param data 数据
     * @param timeout 超时时间
     * @returns 
     */
    public request(route: number, data: any, timeout?: number): Promise<Message> {
        return new Promise((resolve, reject) => {
            if (this.isConnected()) {
                let group = this.waitgroup.get(route);

                if (group === undefined) {
                    group = { seq: 0, callback: new Map() };
                    this.waitgroup.set(route, group);
                }

                let seq = ++group.seq;
                let timeoutId: any;

                if (timeout && timeout > 0) {
                    timeoutId = setTimeout(() => {
                        reject();
                    }, timeout);
                }

                group.callback.set(seq, (message: Message) => {
                    timeoutId && clearTimeout(timeoutId);

                    resolve(message);
                });

                this.send({ seq, route, data });
            } else {
                reject();
            }
        })
    }

    /**
     * 调用回调
     * @param message 消息
     * @returns 回调结果
     */
    private invoke(message: Message): boolean {
        if (message.seq == 0) {
            return false
        }

        const group = this.waitgroup.get(message.route);

        if (group === undefined) {
            return false
        }

        const callback = group.callback.get(message.seq || 0);

        group.callback.delete(message.seq || 0);

        callback && callback(message);

        return true;
    }
}