"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Packer = void 0;
const bytebuffer_1 = __importDefault(require("bytebuffer"));
const json_1 = require("./encoding/json");
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
class Packer {
    constructor(opts) {
        this.opts = opts || { byteOrder: DEFAULT_BYTE_ORDER, routeBytes: DEFAULT_ROUTE_BYTES, seqBytes: DEFAULT_SEQ_BYTES, encoding: new json_1.Json() };
        this.opts.byteOrder = this.opts.byteOrder !== undefined ? this.opts.byteOrder : DEFAULT_BYTE_ORDER;
        this.opts.routeBytes = this.opts.routeBytes !== undefined ? this.opts.routeBytes : DEFAULT_ROUTE_BYTES;
        this.opts.seqBytes = this.opts.seqBytes !== undefined ? this.opts.seqBytes : DEFAULT_SEQ_BYTES;
        this.opts.encoding = this.opts.encoding !== undefined ? this.opts.encoding : new json_1.Json();
        this.buffer = new bytebuffer_1.default(bytebuffer_1.default.DEFAULT_CAPACITY, this.opts.byteOrder != BIG_ENDIAN, bytebuffer_1.default.DEFAULT_NOASSERT);
    }
    /**
     * 打包心跳
     * @returns ArrayBuffer
     */
    packHeartbeat() {
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
    packMessage(message) {
        var _a;
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
        message.data && buffer.append((_a = this.opts.encoding) === null || _a === void 0 ? void 0 : _a.encode(message.data));
        buffer.writeInt32(buffer.offset - DEFAULT_SIZE_BYTES, 0);
        return buffer.flip().toArrayBuffer();
    }
    /**
     * 解包消息
     * @param data 二进制数据
     * @returns Message
     */
    unpack(data) {
        var _a;
        const buffer = this.buffer.clone().append(data, 'binary').flip().skip(DEFAULT_SIZE_BYTES);
        const header = buffer.readUint8();
        const isHeartbeat = header >> 7 == 1;
        if (isHeartbeat) {
            let millisecond;
            if (buffer.remaining() > 0) {
                millisecond = Number(buffer.readInt64().toString());
            }
            return { isHeartbeat, millisecond };
        }
        else {
            const message = { seq: 0, route: 0, data: undefined };
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
                message.data = (_a = this.opts.encoding) === null || _a === void 0 ? void 0 : _a.decode(buffer.readBytes(buffer.remaining()));
            }
            return { isHeartbeat, message };
        }
    }
}
exports.Packer = Packer;
