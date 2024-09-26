"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProtoEncoder = void 0;
class ProtoEncoder {
    encode(data) {
        return data.serializeBinary();
    }
    decode(buff) {
        return buff;
    }
}
exports.ProtoEncoder = ProtoEncoder;
