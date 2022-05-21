package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//条件分支判断语句
func condition() {
	/*
		if 条件1 {
			逻辑代码1
		}else if 条件2{
			逻辑代码2
		}else if ...{
			逻辑代码 ...
		} else {
			逻辑代码 else
		}
	*/
}

// 单分支判断
func singleBranch() {
	score := 88
	if score >= 60 {
		fmt.Println("成绩及格")
	}
}

// 双分支判断
func doubleBranch() {
	score := 88
	if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}
}

//多分支判断
func multipleBranch() {
	score := 88
	if score >= 90 {
		fmt.Println("成绩登记为A")
	} else if score >= 80 {
		fmt.Println("成绩登记为B")
	} else if score >= 70 {
		fmt.Println("成绩登记为C")
	} else if score >= 60 {
		fmt.Println("成绩登记为D")
	} else {
		fmt.Println("成绩登记为E,不及格")
	}
}

//高级条件分支判断语句
func advanceBranch() {
	/*
		if statement;condition{
		}
	*/
	if score := 88; score >= 60 {
		fmt.Println("成绩及格")
	}
}

//选择语句写法
func choose() {
	/*
		switch 表达式{
			case 表达式值1:
				业务逻辑代码1
			case 表达式值2:
				业务逻辑代码2
			case 表达式值3:
				业务逻辑代码3
			case 表达式值...:
				业务逻辑代码...
			default:
				默认逻辑代码
		}
	*/

	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Your score is between 90 and 100.")
	case "B":
		fmt.Println("Your score is between 80 and 90.")
	case "C":
		fmt.Println("Your score is between 70 and 80.")
	case "D":
		fmt.Println("Your score is between 60 and 70.")
	default:
		fmt.Println("Your score is blew 60")
	}
}

//多条件语句选择写法
func multipleCase() {
	month := 5
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("该月份有31天")
	case 4, 6, 9, 11:
		fmt.Println("该月份有30天")
	case 2:
		fmt.Println("该月份闰年为29天，非闰年为28天")
	default:
		fmt.Println("输入有误")
	}
}

//高级选择条件写法
func advanceChoose() {
	/*
		switch statement;expression{
		}
	*/

	switch month := 5; month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("该月份有31天")
	case 4, 6, 9, 11:
		fmt.Println("该月份有30天")
	case 2:
		fmt.Println("该月份闰年为29天，非闰年为28天")
	default:
		fmt.Println("输入有误")
	}
}

func getResult(args ...int) bool {
	for _, v := range args {
		if v < 60 {
			return false
		}
	}
	return true
}

//无表达式的switch
func noCase() {
	score := 88
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("grade A")
	case score >= 80 && score < 90:
		fmt.Println("grade B")
	case score >= 70 && score < 80:
		fmt.Println("grade C")
	case score >= 60 && score < 70:
		fmt.Println("grade D")
	case score < 60:
		fmt.Println("grade E")
	}
}

//fallthrough
func through() {
	s := "从0到Go语言微服务架构师"
	switch {
	case s == "从0到Go语言微服务架构师":
		fmt.Println("从0到Go语言微服务架构师")
		fallthrough
	case s == "Go语言微服务架构核心22讲":
		fmt.Println("Go语言微服务架构核心22讲") //其他语言的break
	case s != "Go语言极简一本通":
		fmt.Println("Go语言极简一本通")
	}
}

//循环语句
func loop() {
	/*
		//三个表达式
		for initialisation;condition;post{
		}
		//一个表达式
		for condition{
		}
		//for 接一个range表达式
		for range_expression{
		}
		//不接表达式,可以用break退出循环
		for{
		}
		for ;;{
		}
		continue
		c语言中，有for,while,do while
	*/
}

//continue语句
func continueLoop() {
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Println(num)
	}
}

//break语句
func breakLoop() {
	num := 0
	for {
		if num > 5 {
			break
		}
		fmt.Println(num)
		num++
	}
}

//一个条件表达式
func singleLoop() {
	num := 0
	for num < 4 {
		fmt.Println(num)
		num++
	}
}

// 三个表达式
func threeExpressionLoop() {
	for num := 0; num < 4; num++ {
		fmt.Println(num)
	}
}

func rangeLoop() {
	str := "从0到Go语言微服务架构师"
	for index, value := range str {
		fmt.Printf("index %d,value %c\n", index, value)
	}
}

//defer延迟调用

func bookPrint() {
	fmt.Println("Go语言极简一本通")
}

type Book struct {
	bookName, authorName string
}

func (b Book) printName() {
	fmt.Printf("%s %s", b.bookName, b.authorName)
}

var s string = "Go语言微服务架构核心22讲"

func showLesson() string {
	defer func() {
		s = "从0到Go语言微服务架构师"
	}()
	fmt.Println("showLesson:s=", s)
	return s
}

func deferClean() {
	/*
		r:=getResource()
		...
		if ...{
			r.release()
			return
		}
		...
		if ...{
			r.release()
			return
		}
		...
		if ...{
			r.release()
			return
		}
		...
		if ...{
			r.release()
			return
		}
		...
		r.release()
		return

		使用defer
		r:=getResource()
		defer r.release()
		if ...{
			...
			return
		}
		if ...{
			...
			return
		}
		if ...{
			...
			return
		}
		if ...{
			...
			return
		}
		...
		return

	*/
}

func visitUrl() {
	url := "http://www.go-edu.cn"
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	if r.StatusCode != http.StatusOK {
		panic(r.StatusCode)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("请求成功...%s\n", body)
	n := rand.Intn(10)
	fmt.Printf("休眠 %d 秒", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.go-edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		panic("http get error" + err.Error())
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("read error" + err.Error())

	}
	fmt.Println(string(body))
	n := rand.Intn(10)
	fmt.Printf("休眠 %d 秒", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func main() {
	url := "http://www.go-edu.cn/"
	for {
		//visitUrl()
		download(url)
	}

	//choose()

	//chinese := 88
	//math := 90
	//english := 95
	//
	//switch getResult(chinese, math, english) {
	//case true:
	//	fmt.Println("考试通过")
	//case false:
	//	fmt.Println("考试未通过")
	//}
	//fmt.Println("--------------------")
	//through()
	//fmt.Println("--------------------")
	//singleLoop()
	//fmt.Println("--------------------")
	//threeExpressionLoop()
	//fmt.Println("--------------------")
	//rangeLoop()
	//fmt.Println("--------------------")
	//continueLoop()
	//fmt.Println("--------------------")
	//breakLoop()

	//延迟调用 函数
	//defer bookPrint()
	//fmt.Println("main...")

	//str := "Go语言极简一本通"
	//defer fmt.Println(str)
	//str = "欢喜"
	//fmt.Println(str)

	//延迟调用方法
	//book := Book{
	//	bookName:   "《Go语言极简一本通》",
	//	authorName: "欢喜",
	//}
	//defer book.printName()
	//fmt.Println("main...")

	//defer栈
	//defer fmt.Println("从0到Go语言微服务架构师")
	//defer fmt.Println("Go语言微服务架构核心22讲")
	//defer fmt.Println("《Go语言极简一本通》")
	//fmt.Println("main...")

	//lesson := showLesson()
	//fmt.Println("main s=", s)
	//fmt.Println("main: lesson=", lesson)

	//goto,goto语句与标签之间不能有变量声明，否则编译报错。
	fmt.Println("从0到Go语言微服务架构师")
	goto label
	fmt.Println("Go语言微服务架构核心22讲")
	//var x int = 0
label:
	fmt.Println("《Go语言极简一本通》")
}
