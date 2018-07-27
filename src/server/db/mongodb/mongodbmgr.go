package mongodbmgr

import (
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 连接消息
var dialContext = new(mongodb.DialContext)

func init() {
	Connect()
}
func Connect() {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer c.Close()
	// index
	c.EnsureUniqueIndex("game", "login", []string{"name"})
	log.Release("mongodb Connect success")
	dialContext = c

	Test()
}
func Test() {
	err :=Find("game","login",bson.M{"name":"hello"})
	if err == nil {
		log.Debug("Test have data,regFail!", )
	}else{
		err =Insert("game","login",bson.M{"name":"hello","password":"123456"})
		if err != nil {
			fmt.Println(err)
			log.Debug("Test write in fail", )
		}
	}
}

func Example() {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// session
	s := c.Ref()
	defer c.UnRef(s)
	err = s.DB("test").C("counters").RemoveId("test")
	if err != nil && err != mgo.ErrNotFound {
		fmt.Println(err)
		return
	}

	// auto increment
	err = c.EnsureCounter("test", "counters", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 3; i++ {
		id, err := c.NextSeq("test", "counters", "test")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
	}

	// index
	c.EnsureUniqueIndex("test", "counters", []string{"key1"})

	// Output:
	// 1
	// 2
	// 3
}

func Find(db string, collection string, docs interface{}) error{
	c:=dialContext
	s := c.Ref()
	defer c.UnRef(s)

	type person struct {
		Id_ bson.ObjectId `bson:"_id"`
		Name string `bson:"name"`
	};
	user:=new(person);
	err := s.DB(db).C(collection).Find(docs).One(&user)
	//idStr:=user.Id_.Hex()
	//fmt.Println(idStr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err;
}

// goroutine safe
func Insert(db string, collection string, docs interface{}) error {
	c:=dialContext
	s := c.Ref()
	defer c.UnRef(s)

	//// 创建索引
	//index := mgo.Index{
	//	Key:        []string{"name"}, // 索引字段， 默认升序,若需降序在字段前加-
	//	Unique:     true,             // 唯一索引 同mysql唯一索引
	//	DropDups:   true,             // 索引重复替换旧文档,Unique为true时失效
	//	Background: true,             // 后台创建索引
	//}
	//if err := s.DB(db).C(collection).EnsureIndex(index); err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//if err := s.DB(db).C(collection).EnsureIndexKey("$2dsphere:location"); err != nil { // 创建一个范围索引
	//	fmt.Println(err)
	//	return err
	//}

	err := s.DB(db).C(collection).Insert(docs)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
