package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"regexp"
	"server/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func init() {
	// 向当前模块（game 模块）注册 消息处理函数
	handler(&msg.Test{}, handleTest)
	handler(&msg.UserLogin{}, handleUserLogin)
	handler(&msg.UserRegister{}, handleUserRegister)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleTest(args []interface{}) {
	// 收到的 Test 消息
	m := args[0].(*msg.Test)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.GetTest())

	retBuf := &msg.Test{
		Test: *proto.String("client"),
	}
	// 给发送者回应一个 Test 消息
	a.WriteMsg(retBuf)
}

func handleUserRegister(args []interface{}) {
	m := args[0].(*msg.UserRegister)
	a := args[1].(gate.Agent)
	name :=m.GetLoginName()
	pwd :=m.GetLoginPW()
	log.Debug("receive UserRegister name=%v", m.GetLoginName())

	reg := regexp.MustCompile(`/^[a-zA-Z\d]\w{2,10}[a-zA-Z\d]$/`)
	matched := reg.FindString(name)
	if(matched!=" "){
		log.Debug("UserRegister name is licit", )
	}
	err := mongodbmgr.Find("game","login",bson.M{"name":name,})
	if err == nil {
		fmt.Println(err)
		log.Debug("UserRegister find in fail", )
		retBuf := &msg.UserResult{
			RetResult: msg.Result_REGISTER_FAIL,
			ErrorInfo:"账号已存在！",
		}
		a.WriteMsg(retBuf)
	}
	err =mongodbmgr.Insert("game","login",bson.M{"name":name,"password":pwd})
	if err != nil {
		fmt.Println(err)
		log.Debug("UserRegister write in fail", )
		retBuf := &msg.UserResult{
			RetResult: msg.Result_REGISTER_FAIL,
			ErrorInfo:"注册失败，请稍后再试！",
		}
		a.WriteMsg(retBuf)
	} else{
		log.Debug("UserRegister write in success", )
		retBuf := &msg.UserResult{
			RetResult: msg.Result_REGISTER_SUCCESS,
		}
		a.WriteMsg(retBuf)
	}
}

func handleUserLogin(args []interface{}) {
	m := args[0].(*msg.UserLogin)
	a := args[1].(gate.Agent)
	log.Debug("receive UserLogin name=%v", m.GetLoginName())

	retBuf := &msg.UserResult{
		RetResult: msg.Result_LOGIN_SUCCESS,
		ErrorInfo:"登陆失败，请稍后再试！",
	}
	a.WriteMsg(retBuf)
}
