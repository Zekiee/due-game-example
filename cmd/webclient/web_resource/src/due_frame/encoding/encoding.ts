export interface Encoding {
    // 编码
    encode(data: any): any
    // 解码
    decode(buff: any): any
}


export class ProtoEncoder implements Encoding {
    public encode(data: any): any {
        return data.serializeBinary();
    }

    public decode(buff: any): any {
        return buff
    }
}
