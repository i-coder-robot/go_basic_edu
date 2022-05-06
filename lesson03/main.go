package main //声明main包
//package name

import (
	"fmt"
	"os"
)

//import name

//导入fmt包，带引字符串时，需要用到它

func main() { //声明main主函数
	/*
			书名：Go语言极简一本通
			公众号：面向加薪学习

		func 函数名(参数列表) (返回值列表){
			//执行体（业务逻辑）
		}
		ln=line
	*/
	fmt.Println("《Go语言极简一本通》")
	fmt.Println("公众号：面向加薪学习")
	os.Getgid()
}
