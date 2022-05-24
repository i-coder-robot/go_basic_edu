package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		func functionName(parameterList){
			//code
		}
		正常调用一个函数
		functionName(parameterList)
		开启一个协程执行这个函数
		go functionName(parameterList)
	*/

	//开启一个协程执行PrintInfo()
	go PrintInfo()
	//使主协程休眠1秒
	time.Sleep(1 * time.Second)
	//打印main
	fmt.Println("main...")

	//启动一号协程
	go PrintNum(1)
	//启动二号协程
	go PrintNum(2)

	time.Sleep(3 * time.Second)
}

func PrintNum(num int) {
	for i := 0; i < 3; i++ {
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintInfo() {
	fmt.Println("从0到Go语言微服务架构师")
}
