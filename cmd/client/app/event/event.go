package event

import (
	"due-game-example/cmd/client/app/message_item"
	"due-game-example/util/logger"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
)

func Init(proxy *client.Proxy) {
	proxy.AddEventListener(cluster.Connect, func(conn *client.Conn) {
		err := conn.Push(message_item.MsgLobbyEchoReq())
		err = conn.Push(message_item.MsgMatchEchoReq())
		err = conn.Push(message_item.MsgRpsGameEchoReq())
		if err != nil {
			logger.Errorf("push message failed: %v", err)
		}
	})

	proxy.AddEventListener(cluster.Reconnect, func(proxy *client.Conn) {
		logger.Errorf("Reconnect")
	})

	proxy.AddEventListener(cluster.Disconnect, func(proxy *client.Conn) {
		logger.Errorf("Disconnect")
	})
}
