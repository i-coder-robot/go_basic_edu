package main

import "fmt"

func main() {
	//通过定义普通变量获取指针
	x := "面向加薪学习"
	ptr := &x
	fmt.Println("x=", x)
	fmt.Println("*ptr=", *ptr)
	fmt.Println("&x=", &x)
	fmt.Println("ptr=", ptr)
	fmt.Println("------------------")
	//new先创建指针并分配好内存，再给指针写入值
	ptr2 := new(string)
	*ptr2 = "从0到Go语言微服务架构师"
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	//先声明一个指针变量，再从其他变量获取内存地址指针变量
	x2 := "Go语言微服务架构核心22讲"
	var p *string
	p = &x2
	fmt.Println(p)

	//& 可以从一个变量中取得内存地址
	//* 赋值的左边，指该指针指向的变量， 赋值的右边-指总一个指针变量中取得变量值（解引用）

	fmt.Println("------------------")
	pointerType()
	fmt.Println("------------------")
	zeroPointer()
	fmt.Println("------------------")

	x3 := 99
	p3 := &x3
	fmt.Println("执行changeByPointer函数之前p3是", *p3)
	changeByPointer(p3)
	fmt.Println("执行changeByPointer函数之后p3是", *p3)

	fmt.Println("------------------")
	x4 := [3]int{10, 20, 30}
	ChangeSlice(x4[:])
	fmt.Println(x4)

	y := [3]int{10, 20, 30}
	ChangeArray(&y)
	fmt.Println(y)

	//x5 := [...]int{10, 20, 30}
	//ptr5 := &x5
	//ptr5++ //会报错，不支持指针的运算
}

// ChangeSlice 修改指针
func ChangeSlice(value []int) {
	value[0] = 200
}

// ChangeArray 修改数组指针
func ChangeArray(value *[3]int) {
	(*value)[0] = 200
}

func changeByPointer(value *int) {
	*value = 200
}

func zeroPointer() {
	x := "从0到Go语言微服务架构师"
	var ptr *string
	fmt.Println("ptr is ", ptr)
	ptr = &x
	fmt.Println("ptr is ", ptr)
}

func pointerType() {
	mystr := "Go语言极简一本通"
	myint := 1
	mybool := false
	myfloat := 3.2
	fmt.Printf("type of &mystr is :%T\n", &mystr)
	fmt.Printf("type of &myint is :%T\n", &myint)
	fmt.Printf("type of &mybool is :%T\n", &mybool)
	fmt.Printf("type of &myfloat is :%T\n", &myfloat)
}
