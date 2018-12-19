package main

import (
	"os/exec"
	"fmt"
	"strings"

)

func main(){
	StartService()
}

func StartService(){
	start("nsqlookupd &")
	start("nsqd --lookupd-tcp-address=127.0.0.1:4160 &")
	start("nsqadmin --lookupd-http-address=127.0.0.1:4161 &")
	start("nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161")
	fmt.Println("success")
}

func StopService(){

}

func start(commands string){
	fmt.Println(commands)  // 输出正在执行的命令
	args := strings.Split(commands, " ")   // 以空格分割字符串
	service := exec.Command(args[0], args[1:]...)
	// 
	service.Start()
	fmt.Println(args[0], " pid is ", service.Process.Pid)
	// service.Process.Pid 获取service 进程的pid

	/*
	if err := service.Start(); err != nil{
		fmt.Println(err)
		return 
	}
	*/
}
