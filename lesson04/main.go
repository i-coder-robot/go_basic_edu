package main

import (
	"fmt"
)

//const (
//	BEIJING  = 0
//	SHANGHAI = 1
//	SHENZHEN = 2
//)

const (
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

func main() {
	//var关键字
	//方法1 声明一个变量，默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("a的类型是:%T\n", a)

	//方法2 声明一个变量，并且初始化一个值
	var b int = 100
	fmt.Printf("b=%d,b的类型是:%T\n", b, b)

	//方法3 初始化的时候，去掉数据类型,Go语言通过值自动匹配类型
	var c = 100
	fmt.Printf("c=%d,c的类型是:%T\n", c, c)

	var cc = "Go语言微服务架构师核心22讲"
	fmt.Printf("cc=%s,cc的类型是:%T\n", cc, cc)

	//方法4 短声明 :=

	e := 100
	fmt.Printf("e=%d,e的类型是:%T\n", e, e)

	f := "Go语言极简一本通"
	fmt.Printf("f=%s,f的类型是%T\n", f, f)

	//var xx, yy int = 100, 200
	//var kk, wx = 300, "write_code_666（欢喜哥）"
	//
	//var (
	//	nn     int = 100
	//	wechat     = "write_code_666（欢喜哥）"
	//)

	const length = 10
	//length = 20
	fmt.Println("length=", length)

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)

	//iota常量计数器

}
