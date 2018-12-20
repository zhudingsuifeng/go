package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer 

// nsqd 监听的ip端口
var nsqip = "127.0.0.1:4150"

// 初始化生产者
func InitProducer(ip string) {
	var err error
	// fmt.Println("address: ", ip)
	producer, err = nsq.NewProducer(ip, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

// 发布消息
func SendMessage(topic string, message string) error {
	var err error
	InitProducer(nsqip)
	if producer != nil{
		// 生产者创建成功
		if message == "" {
			// 不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) //发布消息
		fmt.Println("send ", message, "on ", topic)
		producer.Stop()    // 关闭
		return err
	} 
	return nil
	// return fmt.Errorf("producer is nil", err)
}

/*
// 推送消息
func Send(message string) {
	strIP1 := "127.0.0.1:4150"
	InitProducer(strIP1)
	running := true
	// 读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop"{
			running = false
		}
		for err := SendMessage("test", command); err != nil; err = SendMessage("test", command){

		}
	}
}
*/