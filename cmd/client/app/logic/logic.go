package logic

import (
	"due-game-example/shared/pb/pb_common"
	"due-game-example/shared/pb/pb_lobby"
	"due-game-example/util/logger"
	"github.com/dobyte/due/v2/cluster/client"
)

type logic struct {
	proxy *client.Proxy
}

func NewLogic(proxy *client.Proxy) *logic {
	return &logic{proxy: proxy}
}

func (l *logic) Init() {
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_LOBBY_ECHO), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_LOBBY_LOGIN), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_MATCH_ECHO), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_MATCH_JOIN), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_RPSGAME_ECHO), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_RPSGAME_START), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_RPSGAME_PLAY), l.echo)
	l.proxy.AddRouteHandler(int32(pb_common.ROUTE_RPSGAME_RESULT), l.echo)
}

func (l *logic) echo(r *client.Context) {
	res := &pb_lobby.EchoRes{}

	err := r.Parse(res)
	if err != nil {
		logger.Errorf("invalid login message, err: %v", err)
		return
	}

	logger.Info(res)
}
