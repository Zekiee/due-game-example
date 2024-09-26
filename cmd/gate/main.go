package main

import (
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := gate.NewGate(
		gate.WithServer(ws.NewServer()),
		gate.WithLocator(redis.NewLocator()),
		gate.WithRegistry(consul.NewRegistry()),
	)
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
