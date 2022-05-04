package main

import "fmt"

//声明枚举值
//const (
//	BEIJING  = 0
//	SHANGHAI = 1
//	SHENZHEN = 2
//)

//const (
//	BEIJING = iota
//	SHANGHAI
//	SHENZHEN
//)

const (
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

func main() {
	//方法一 声明一个变量，默认值0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)
	//方法二 声明一个变量，初始化一个值
	var b int = 100
	fmt.Printf("b=%d, type of b = %T\n", b, b)
	//方法三 初始化的时候，省去类型，通过值自动识别当前类型。
	var book string = "《Go语言极简一本通》"
	fmt.Printf("book=%s, type of book = %T\n", book, book)
	//方法四 常用方法，省去var 使用:=形式声明，推导数据类型，并赋值。
	//注意：短声明在函数或方法内部使用，不支持全局声明。
	var c = 100
	fmt.Printf("c=%d, type of c = %T\n", c, c)

	var name = "公众号：面向加薪学习"
	fmt.Printf("name=%s, type of name = %T\n", name, name)

	d := 200
	fmt.Printf("d=%d, type of d = %T\n", d, d)

	name2 := "欢喜"
	fmt.Printf("name2=%s, type of name2 = %T\n", name2, name2)

	// 声明多个变量
	//var (
	//	xx, yy int = 100, 300
	//)
	//var(
	//	nn int =100
	//	mm bool = true
	//)

	const length = 10
	//声明的常量，不允许修改
	//length = 20
	fmt.Println("length=", length)

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)

	//iota 常量计数器

	//var func = 0
}
