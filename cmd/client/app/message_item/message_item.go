package message_item

import (
	"due-game-example/shared/pb/pb_common"
	"due-game-example/shared/pb/pb_lobby"
	"due-game-example/shared/pb/pb_match"
	"due-game-example/shared/pb/pb_rpsgame"
	"github.com/dobyte/due/v2/cluster"
)

func MsgLobbyEchoReq() *cluster.Message {
	return &cluster.Message{
		Route: int32(pb_common.ROUTE_LOBBY_ECHO),
		Data: &pb_lobby.EchoReq{
			Msg: "Hello there! pb_lobby",
		},
	}
}

func MsgMatchEchoReq() *cluster.Message {
	return &cluster.Message{
		Route: int32(pb_common.ROUTE_MATCH_ECHO),
		Data: &pb_match.EchoReq{
			Msg: "Hello there! pb_match",
		},
	}
}

func MsgRpsGameEchoReq() *cluster.Message {
	return &cluster.Message{
		Route: int32(pb_common.ROUTE_RPSGAME_ECHO),
		Data: &pb_rpsgame.EchoReq{
			Msg: "Hello there! pb_rpsgame",
		},
	}
}
