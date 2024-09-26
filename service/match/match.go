package match

import (
	"due-game-example/shared/pb/pb_common"
	"due-game-example/util/tools"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/transport/rpcx/v2"
	"github.com/dobyte/due/v2/cluster/node"
)

func NewMatchNode() *node.Node {
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := consul.NewRegistry()
	node := node.NewNode(
		node.WithName("match"),
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(rpcx.NewTransporter()),
	)
	InitRoute(node)
	return node
}

func InitRoute(n *node.Node) {
	l := NewMatchService()
	// 注册路由
	g := n.Proxy().Router().Group()
	g.AddRouteHandler(int32(pb_common.ROUTE_MATCH_ECHO), false, tools.Wrap(l.EchoHandler))

}
