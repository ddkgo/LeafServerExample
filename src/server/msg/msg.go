package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)
// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&Hello{})
}
