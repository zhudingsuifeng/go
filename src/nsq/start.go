package main

import (
	"fmt"
	"strings"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = os.Args[1]
	app.Usage = "send messages used nsq"
	app.Action = func (c *cli.Context) error {
		switch os.Args[1] {
		case "start":
			Start()
		case "send":
			err := SendMessage(os.Args[2], os.Args[3])
			if err != nil{
				log.Fatal(err)
			}
		case "stop":
			Stop() // 暂时还没有实现
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}

// 启动nsq相关的服务程序
func Start(){
	Service("nsqlookupd &")
	Service("nsqd --lookupd-tcp-address=127.0.0.1:4160 &")
	Service("nsqadmin --lookupd-http-address=127.0.0.1:4161 &")
	Service("nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161")
	fmt.Println("nsq start up ...")
}

// 具体要启动的服务
func Service(Args string){
	fmt.Println(Args)  // 输出正在执行的命令
	args := strings.Split(Args, " ")   // 以空格分割字符串
	command := exec.Command(args[0], args[1:]...)
	// 执行外部程序
	command.Start()
	// fmt.Println(args[0], " pid is ", service.Process.Pid)
	// service.Process.Pid 获取service 进程的pid

	/*
	if err := service.Start(); err != nil{
		fmt.Println(err)
		return 
	}
	*/
}

// 停止nsq相关服务
func Stop(){
	fmt.Println("暂时还未实现")
}
