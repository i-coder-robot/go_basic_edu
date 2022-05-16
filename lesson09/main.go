package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		func function_name(parameter_list) (result_list){
			//函数体
		}

		func (parameter_list) (result_list){
				//函数体
			}

		函数可见性
		1 函数首字母大写,对于所有包是public，其他包任意调用
		2 函数首字母小写，这个函数是private，其他包无法访问。
	*/

	fmt.Println("56+1=", sum(56, 1))
	printBookName()
	fmt.Println("58-2=", sub(58, 2))
	fmt.Println("------------------")
	fmt.Println(show("Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"))
	fmt.Println("------------------")
	PrintType(57, 3.14, "从0到Go语言微服务架构师")
	fmt.Println("------------------")
	var s []string
	s = append(s, []string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}...)
	fmt.Println(s)
	fmt.Println("------------------")
	bookInfo, err := showBookInfo("《Go语言极简一本通》", "欢喜")
	fmt.Printf("bookInfo=%s,err=%v\n", bookInfo, err)
	bookInfo, err = showBookInfo("《Go语言极简一本通》", "")
	fmt.Printf("bookInfo=%s,err=%v\n", bookInfo, err)
}

func showBookInfo(bookName, authorName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称不能为空")
	}
	if authorName == "" {
		return "", errors.New("作者名称不能为空")
	}
	return bookName + ",作者:" + authorName, nil
}

func showBookInfo2(bookName, authorName string) (info string, err error) {
	if bookName == "" {
		err = errors.New("图书名称不能为空")
		return
	}
	if authorName == "" {
		err = errors.New("作者名称不能为空")
		return
	}
	//不适用:=因为已经在返回值那里声明了
	info = bookName + ",作者:" + authorName
	//直接返回即可
	return
}

func PrintType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int")
		case string:
			fmt.Println(arg, "type is string")
		case float64:
			fmt.Println(arg, "type is float64")
		default:
			fmt.Println(arg, "is an  unknown type")
		}
	}
}

func show(args ...string) int {
	sum := 0
	for _, item := range args {
		fmt.Println(item)
		sum += 1
	}
	return sum
}

//函数返回一个无命名变量，返回值列表的括号省略。
func sum(x int, y int) int {
	return x + y
}

//无参数列表和返回值
func printBookName() {
	fmt.Println("《Go语言极简一本通》")
}

//参数类型一致，只在最后一个参数后添加类型
func sub(x, y int) int {
	return x - y
}
