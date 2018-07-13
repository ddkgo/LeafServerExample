package gate

import (
	"server/msg"
	"server/game"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
