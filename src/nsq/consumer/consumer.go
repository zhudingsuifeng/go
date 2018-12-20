package main

import (
	"fmt"
	"time"
	"github.com/nsqio/go-nsq"
)

// 消费者
// type ConsumerT struct{}

// 主函数
func main(){
	// 初始化消费者(消息接收端)
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	fmt.Println("I am listeming...")
	running := true
	for running {
		time.Sleep(time.Second * 10)
		/*
		if command == "stop"{
			running = false
		}
		*/
	}
}

/*
// 消费者处理消息的另一种方法
func (*ConsumerT) HandleMessage(msg *nsq.Message) error{
	fmt.Println("receive", msg.NSQDAddress, "message at ", time.Now().Format("2006-01-02 15:04:05"),string(msg.Body))
	return nil
}
*/

func InitConsumer(topic string, channel string, address string){
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second      //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg)    //新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)          // 屏蔽系统日志
	// c.AddHandler(&ConsumerT{})    // 添加消费者接口

	// 当消费者每接收到一则消息，触发下列函数
	c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		// 设置消息处理函数
		fmt.Println("receive", msg.NSQDAddress, "message at ", time.Now().Format("2006-01-02 15:04:05"),string(msg.Body))
		return nil
	}))

	// 建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}
}