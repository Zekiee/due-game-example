"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Json = void 0;
class Json {
    // 编码
    encode(data) {
        return JSON.stringify(data);
    }
    // 解码
    decode(buff) {
        console.log(buff);
        return JSON.parse(buff.toUTF8());
    }
}
exports.Json = Json;
