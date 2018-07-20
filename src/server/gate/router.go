package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// 这里指定消息 路由到 game 模块
	msg.Processor.SetRouter(&msg.Test{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserLogin{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserRegister{}, game.ChanRPC)
}
