package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		定义函数
		function functionName(parameterList){
			//业务逻辑代码
		}
		调用函数
		functionName(parameterList)
		开启一个协程调用函数
		go functionName(parameterList)
	*/

	//go PrintInfo()
	//
	//time.Sleep(1 * time.Second)
	//
	//fmt.Println("main...")

	//开启1号协程
	go PrintNum(1)

	//开启2号协程
	go PrintNum(2)

	time.Sleep(time.Second)
}

func PrintNum(num int) {
	for i := 0; i < 3; i++ {
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintInfo() {
	fmt.Println("从0到GO语言微服务架构师")
}
