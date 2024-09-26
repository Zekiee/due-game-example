import { Encoding } from "./encoding";

export class Json implements Encoding  {
    // 编码
    public encode(data: any): any {
        return JSON.stringify(data);
    }

    // 解码
    public decode(buff: any): any {
        console.log(buff)
        return JSON.parse(buff.toUTF8());
    }
}