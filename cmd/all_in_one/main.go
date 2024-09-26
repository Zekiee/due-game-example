package main

import (
	"due-game-example/service/lobby"
	"due-game-example/service/match"
	"due-game-example/service/rpsgame"
	"due-game-example/shared/pb/pb_lobby"
	"github.com/dobyte/due/v2"
	"github.com/golang/protobuf/proto"
)

func main() {
	// 创建容器
	req := pb_lobby.EchoReq{Msg: "hello"}
	b, e := proto.Marshal(&req)
	sb := string(b)
	_ = e
	_ = sb
	container := due.NewContainer()
	// 添加lobby
	container.Add(lobby.NewLobbyNode())
	// 添加match
	container.Add(match.NewMatchNode())
	// 添加RPSGame
	container.Add(rpsgame.NewRPSGameNode())
	// 启动容器
	container.Serve()
}
