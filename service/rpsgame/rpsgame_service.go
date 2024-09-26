package rpsgame

import (
	"context"
	"due-game-example/shared/pb/pb_common"
	"due-game-example/shared/pb/pb_lobby"
	"due-game-example/shared/pb/pb_rpsgame"
	"due-game-example/util/logger"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/session"
	"github.com/dobyte/due/v2/utils/xcall"
	"time"
)

type RPSGameService struct {
	proxy   *node.Proxy
	msgChan chan *pb_rpsgame.BotStatusReq
}

func NewRPSGameService(p *node.Proxy) *RPSGameService {
	r := &RPSGameService{
		proxy:   p,
		msgChan: make(chan *pb_rpsgame.BotStatusReq, 1000),
	}

	xcall.Go(r.processMsg)

	return r
}

// 处理消息队列中的消息
func (l *RPSGameService) processMsg() {
	ticker := time.NewTicker(25 * time.Millisecond) // 定时器，每隔100ms触发
	defer ticker.Stop()

	res := &pb_rpsgame.BotStatusRes{}

	for {
		select {
		case msg := <-l.msgChan:
			// 直接收集消息，无需等待ticker
			res.BotStatus = append(res.BotStatus, msg)

		case <-ticker.C:
			// 定时检查并发送消息
			if len(res.BotStatus) > 0 {
				l.sendMsg(res)
				res = &pb_rpsgame.BotStatusRes{} // 重置，准备下一轮消息
			}
		}
	}
}

// 发送消息到客户端
func (l *RPSGameService) sendMsg(res *pb_rpsgame.BotStatusRes) {
	logger.Debug("Sending message to client")
	// 在这里添加发送到客户端的逻辑
	data := &cluster.BroadcastArgs{
		Kind: session.Conn,
		Message: &cluster.Message{
			Seq:   0,
			Route: int32(pb_common.ROUTE_RPSGAME_START),
			Data:  res,
		},
	}
	l.proxy.Broadcast(context.Background(), data)
}

func (l *RPSGameService) EchoHandler(ctx node.Context, req *pb_lobby.EchoReq) (res *pb_lobby.EchoRes) {
	logger.Debug("RPSGameService EchoHandler")
	res = &pb_lobby.EchoRes{
		Uid: 1,
		Msg: req.Msg,
	}
	time.Sleep(time.Millisecond * 100)
	return
}

func (l *RPSGameService) RpsgameStart(ctx node.Context, req *pb_rpsgame.BotStatusReq) (res *pb_rpsgame.BotStatusRes) {
	logger.Debug("RPSGameService EchoHandler")
	l.msgChan <- req
	return
}
