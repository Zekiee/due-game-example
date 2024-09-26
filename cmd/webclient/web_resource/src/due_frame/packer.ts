import ByteBuffer from 'bytebuffer';
import { Encoding } from "./encoding/encoding";
import { Json } from "./encoding/json";

export interface PackerOptions {
    // 字节序；默认为big
    byteOrder?: string;
    // 序列号字节长度（字节），长度为0时不开启序列号编码；默认为2字节，最大值为65535
    seqBytes?: number;
    // 路由字节长度（字节）；默认为2字节，最大值为65535
    routeBytes?: number;
    // 编辑解码器
    encoding?: Encoding;
}

export interface Message {
    // 消息序列号
    seq?: number;
    // 消息路由
    route: number;
    // 消息数据
    data?: any;
}

export interface Packet {
    // 是否是心跳包
    isHeartbeat: boolean;
    // 心跳包携带的服务器时间（毫秒）
    millisecond?: number;
    // 消息数据
    message?: Message;
}

// 大端序
const BIG_ENDIAN = 'big';

// 小端序
const LITTLE_ENDIAN = 'little';

// 默认字节序
const DEFAULT_BYTE_ORDER = BIG_ENDIAN;

// 默认size位字节长度
const DEFAULT_SIZE_BYTES = 4;

// 默认header位字节长度
const DEFAULT_HEADER_BYTES = 1;

// 默认route位字节长度
const DEFAULT_ROUTE_BYTES = 2;

// 默认seq位字节长度
const DEFAULT_SEQ_BYTES = 2;

// 默认buffer数据位字节长度
const DEFAULT_BUFFER_BYTES = 5000;

export class Packer {
    private opts: PackerOptions;
    private buffer: any;

    public constructor(opts?: PackerOptions) {
        this.opts = opts || { byteOrder: DEFAULT_BYTE_ORDER, routeBytes: DEFAULT_ROUTE_BYTES, seqBytes: DEFAULT_SEQ_BYTES, encoding: new Json() };
        this.opts.byteOrder = this.opts.byteOrder !== undefined ? this.opts.byteOrder : DEFAULT_BYTE_ORDER;
        this.opts.routeBytes = this.opts.routeBytes !== undefined ? this.opts.routeBytes : DEFAULT_ROUTE_BYTES;
        this.opts.seqBytes = this.opts.seqBytes !== undefined ? this.opts.seqBytes : DEFAULT_SEQ_BYTES;
        this.opts.encoding = this.opts.encoding !== undefined ? this.opts.encoding : new Json();
        this.buffer = new ByteBuffer(ByteBuffer.DEFAULT_CAPACITY, this.opts.byteOrder != BIG_ENDIAN, ByteBuffer.DEFAULT_NOASSERT);
    }

    /**
     * 打包心跳
     * @returns ArrayBuffer
     */
    public packHeartbeat(): ArrayBuffer {
        let buffer = this.buffer.clone();
        let header = 1 << 7;

        buffer.writeInt32(DEFAULT_HEADER_BYTES);

        buffer.writeInt8(header);

        return buffer.flip().toArrayBuffer();
    }

    /**
     * 打包消息
     * @param message 消息数据
     * @returns ArrayBuffer
     */
    public packMessage(message: Message): ArrayBuffer {
        let buffer = this.buffer.clone();
        let header = 0;
        let seq = message.seq || 0;
        let route = message.route || 0;

        buffer.skip(DEFAULT_SIZE_BYTES);

        buffer.writeInt8(header);

        switch (this.opts.routeBytes) {
            case 1:
                buffer.writeInt8(route);
                break;
            case 2:
                buffer.writeInt16(route);
                break;
            case 4:
                buffer.writeInt32(route);
                break;
        }

        switch (this.opts.seqBytes) {
            case 1:
                buffer.writeInt8(seq);
                break;
            case 2:
                buffer.writeInt16(seq);
                break;
            case 4:
                buffer.writeInt32(seq);
                break;
        }

        message.data && buffer.append(this.opts.encoding?.encode(message.data));

        buffer.writeInt32(buffer.offset - DEFAULT_SIZE_BYTES, 0);

        return buffer.flip().toArrayBuffer();
    }

    /**
     * 解包消息
     * @param data 二进制数据
     * @returns Message
     */
    public unpack(data: any): Packet {
        const buffer = this.buffer.clone().append(data, 'binary').flip().skip(DEFAULT_SIZE_BYTES);
        const header = buffer.readUint8();
        const isHeartbeat = header >> 7 == 1;

        if (isHeartbeat) {
            let millisecond

            if (buffer.remaining() > 0) {
                millisecond = Number(buffer.readInt64().toString());
            }

            return { isHeartbeat, millisecond };
        } else {
            const message = { seq: 0, route: 0, data: undefined }

            if (this.opts.routeBytes) {
                if (this.opts.routeBytes > buffer.remaining()) {
                    return { isHeartbeat };
                }

                switch (this.opts.routeBytes) {
                    case 1:
                        message.route = buffer.readInt8();
                        break;
                    case 2:
                        message.route = buffer.readInt16();
                        break;
                    case 4:
                        message.route = buffer.readInt32();
                        break;
                }
            }

            if (this.opts.seqBytes) {
                if (this.opts.seqBytes > buffer.remaining()) {
                    return { isHeartbeat };
                }

                switch (this.opts.seqBytes) {
                    case 1:
                        message.seq = buffer.readInt8();
                        break;
                    case 2:
                        message.seq = buffer.readInt16();
                        break;
                    case 4:
                        message.seq = buffer.readInt32();
                        break;
                }
            }

            if (buffer.remaining() > 0) {                
                message.data = this.opts.encoding?.decode(buffer.readBytes(buffer.remaining()));
            }

            return { isHeartbeat, message };
        }
    }
}