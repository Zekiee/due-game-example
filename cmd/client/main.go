package main

import (
	"due-game-example/cmd/client/app/event"
	"due-game-example/cmd/client/app/logic"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)

	// 初始化事件和路由
	event.Init(component.Proxy())
	logic.NewLogic(component.Proxy()).Init()

	// 初始化监听
	initListen(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化监听
func initListen(proxy *client.Proxy) {
	// 监听组件启动
	proxy.AddHookListener(cluster.Start, startHandler)
}

// 组件启动处理器
func startHandler(proxy *client.Proxy) {
	if _, err := proxy.Dial(); err != nil {
		log.Errorf("gate connect failed: %v", err)
		return
	}
}
