package main

import (
	"fmt"
	"go_basic_edu/lesson10/book"
	//"crypto/rand"
	//mrand "math/rand"
)

//package packageName

//import path
//import(
// path1
// path2
// ...
//)

func init() {
	fmt.Println("init...")
}

func main() {
	/*
		包的初始化顺序
		1 包级别的变量
		2 init() 根据编译器解析的顺序进行调用

		包的匿名导入
		import _ 包名
	*/

	fmt.Println("《Go语言极简一本通》")
	info, _ := book.ShowBookInfo("《Go语言极简一本通》", "欢喜")
	fmt.Println(info)
}
