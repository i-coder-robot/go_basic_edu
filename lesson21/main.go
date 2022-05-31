package main

import "fmt"

func main() {
	//静态类型 变量声明时候的类型
	//var number int
	//var name string

	//动态类型 程序运行时系统才能看见的类型。
	//var in interface{}
	//in = 100
	//in = "《从0到Go语言微服务架构师》"

	var number int = 100
	fmt.Println(number)
	number2 := (int)(100)
	fmt.Println(number2)
	fmt.Printf("number2 type : %T,data:%v\n", number2, number2)
	number3 := (interface{})(100)
	fmt.Println(number3)
	fmt.Printf("number3 type : %T,data:%v\n", number3, number3)
}
