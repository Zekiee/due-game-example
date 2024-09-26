"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Client = void 0;
const bytebuffer_1 = __importDefault(require("bytebuffer"));
const packer_1 = require("./packer");
class Client {
    constructor(opts) {
        this.opts = opts;
        this.websocket = undefined;
        this.packer = opts.packer || new packer_1.Packer(opts.packer);
        this.waitgroup = new Map();
        this.buffer = new bytebuffer_1.default(bytebuffer_1.default.DEFAULT_CAPACITY, false, bytebuffer_1.default.DEFAULT_NOASSERT);
    }
    /**
     * 连接服务区
     * @returns boolean
     */
    connect() {
        try {
            this.disconnect();
            // 新建WS客户端
            this.websocket = new WebSocket(this.opts.url);
            this.websocket.binaryType = 'arraybuffer';
            // 监听WS连接打开
            this.websocket.onopen = (ev) => {
                this.heartbeat();
                this.connectHandler && this.connectHandler();
            };
            // 监听WS连接关闭
            this.websocket.onclose = (ev) => {
                this.disconnectHandler && this.disconnectHandler();
            };
            // 监听WS连接错误
            this.websocket.onerror = (ev) => {
                this.errorHandler && this.errorHandler();
            };
            // 监听WS消息
            this.websocket.onmessage = (e) => {
                if (e.data.byteLength == 0) {
                    return;
                }
                const packet = this.packer.unpack(e.data);
                if (packet.isHeartbeat) {
                    this.heartbeatHandler && this.heartbeatHandler(packet.millisecond);
                }
                else {
                    packet.message && (this.invoke(packet.message) || this.receiveHandler && this.receiveHandler(packet.message));
                }
            };
            return true;
        }
        catch (error) {
            console.log(error);
            return false;
        }
    }
    /**
     * 断开连接
     * @returns void
     */
    disconnect() {
        if (!this.websocket) {
            return;
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
    heartbeat() {
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
    onConnect(handler) {
        this.connectHandler = handler;
    }
    /**
     * 设置断开连接处理器
     * @param handler DisconnectHandler
     */
    onDisconnect(handler) {
        this.disconnectHandler = handler;
    }
    /**
     * 设置消息接收处理器
     * @param handler ReceiveHandler
     */
    onReceive(handler) {
        this.receiveHandler = handler;
    }
    /**
     * 设置错误处理器
     * @param handler ErrorHandler
     */
    onError(handler) {
        this.errorHandler = handler;
    }
    /**
     * 设置心跳处理器
     * @param handler HeartbeatHandler
     */
    onHeartbeat(handler) {
        this.heartbeatHandler = handler;
    }
    /**
     * 检测客户端是否已连接
     * @returns boolean
     */
    isConnected() {
        return this.websocket !== undefined && this.websocket.readyState == WebSocket.OPEN;
    }
    /**
     * 检测客户端是否正在连接
     * @returns boolean
     */
    isConnecting() {
        return this.websocket !== undefined && this.websocket.readyState === WebSocket.CONNECTING;
    }
    /**
     * 发送消息
     * @param message Message 消息
     * @returns boolean
     */
    send(message) {
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
    request(route, data, timeout) {
        return new Promise((resolve, reject) => {
            if (this.isConnected()) {
                let group = this.waitgroup.get(route);
                if (group === undefined) {
                    group = { seq: 0, callback: new Map() };
                    this.waitgroup.set(route, group);
                }
                let seq = ++group.seq;
                let timeoutId;
                if (timeout && timeout > 0) {
                    timeoutId = setTimeout(() => {
                        reject();
                    }, timeout);
                }
                group.callback.set(seq, (message) => {
                    timeoutId && clearTimeout(timeoutId);
                    resolve(message);
                });
                this.send({ seq, route, data });
            }
            else {
                reject();
            }
        });
    }
    /**
     * 调用回调
     * @param message 消息
     * @returns 回调结果
     */
    invoke(message) {
        if (message.seq == 0) {
            return false;
        }
        const group = this.waitgroup.get(message.route);
        if (group === undefined) {
            return false;
        }
        const callback = group.callback.get(message.seq || 0);
        group.callback.delete(message.seq || 0);
        callback && callback(message);
        return true;
    }
}
exports.Client = Client;
