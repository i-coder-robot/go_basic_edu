package main

import "fmt"

func main() {
	//bookFunc := func() {
	//	fmt.Println("《Go语言极简一本通》")
	//}
	//bookFunc()
	//fmt.Printf("bookFunc 的类型是%T\n", bookFunc)

	//f := func(x, y string) string {
	//	return x + y
	//}
	//printRes(f)

	//s := show()
	//fmt.Println(s("欢喜", "《Go语言极简一本通》"))

	//闭包
	x := 100
	func() {
		fmt.Println(x)
	}()
}

//高阶函数
func printRes(show func(author, book string) string) {
	fmt.Println(show("欢喜", "《Go语言极简一本通》"))
}

//高阶函数
func show() func(author, book string) string {
	return func(x, y string) string {
		return x + y
	}
}
