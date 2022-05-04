package main // 声明main包
import (
	"fmt"
	"os"
) // 导入fmt包，打印字符串时候需要用到的。

func main() {
	/*
		一下代码都是将信息打印到屏幕上。
		公众号：面向加薪学习
		func 函数名(参数列表) 返回值列表 {
			代码块（业务逻辑）
		}

		ln = line
	*/
	fmt.Println("以下是我的微信公众号：")
	fmt.Println("《面向加薪学习》")
	os.Getgid()
}
