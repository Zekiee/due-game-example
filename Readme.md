
# due-game-example

## due框架的游戏示例
包含功能
- [ ] 客户端账号登录
- [ ] 匹配, 创建, 加入房间
- [ ] 剪刀石头布游戏

主要涉及
- gate示例
- node示例
- client示例
- admin示例
- mesh示例
- mysql使用
- consul使用
- redis使用
- node内的 routeGroup 示例
- node内的 route wrap 示例 @感谢群友提供
- node内的 middleware 示例
- 使用一个main.go启动多个服务，用于开发环境


## proto编译环境

```shell
// protoc-gen-go v1.25.0
// protoc        v3.20.3
// protoc-gen-gofast latest 
// go install github.com/gogo/protobuf/protoc-gen-gofast@latest
```

protoc
> https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.3

protoc-gen-go
> https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.23.0

protoc-gen-rpcx need build from source code
> https://github.com/rpcxio/protoc-gen-rpcx/releases/tag/v0.2.2