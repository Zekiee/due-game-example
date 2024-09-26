package lobby

import (
	"due-game-example/shared/pb/pb_lobby"
	"due-game-example/util/logger"
	"github.com/dobyte/due/v2/cluster/node"
	"time"
)

type LobbyService struct {
}

func NewLobbyService() *LobbyService {
	return &LobbyService{}
}

func (l *LobbyService) EchoHandler(ctx node.Context, req *pb_lobby.EchoReq) (res *pb_lobby.EchoRes) {
	logger.Debug("LobbyService EchoHandler")
	res = &pb_lobby.EchoRes{
		Uid: 1,
		Msg: req.Msg,
	}
	time.Sleep(time.Millisecond * 100)
	return
}
