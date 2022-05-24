package main

import "fmt"

func main() {
	/*
		type interface_name interface{
			method1()
			method2()
			...
		}
	*/

	student1 := Student{
		name: "张三",
		book: "《Go语言极简一本通通》",
	}
	student1.learn()
	fmt.Println("------------------------")

	var s1 Study
	var s2 Study
	student2 := Student{
		name: "李四",
		book: "Go语言微服务架构核心22讲",
	}
	s1 = student2
	s1.learn()
	student3 := Student{
		name: "王五",
		book: "Go语言微服务架构核心22讲",
	}
	s1 = &student3
	s1.learn()

	worker1 := Worker{
		name: "老王",
		book: "从0到Go语言微服务架构师",
		by:   "视频+答疑+群组",
	}
	s2 = &worker1
	s2.learn()

	//s2=worker1  无法将 'worker1' (类型 Worker) 用作类型 Study 类型未实现 'Study'，因为 'learn' 方法有指针接收器
	fmt.Println("------------------------")
	var s Study
	s = student2
	ShowInterface(s)
	s.learn()

	fmt.Println("------------------------")
	str := "从0到Go语言微服务架构师"
	ShowType(str)
	num := 3.14
	ShowType(num)

	fmt.Println("------------------------")
	var i interface{}
	fmt.Printf("类型:%T,值:%v\n", i, i) //类型:<nil>,值:<nil>

	fmt.Println("------------------------")

	i = "从0到Go语言微服务架构师"
	fmt.Println(i)
	i = 3.14
	fmt.Println(i)

	fmt.Println("------------------------")
	x := make([]interface{}, 3)
	x[0] = "从0到Go语言微服务架构师"
	x[1] = 3.14
	x[2] = []int{1, 2, 3}
	for _, item := range x {
		fmt.Println(item)
	}

	//var num2 = 10
	//var i2 interface{} = num2
	//var str string = i2 //无法将 'i2' (类型 interface{}) 用作类型 string

	//var s3 = []int{1, 2, 3}
	//var i3 interface{} = s3
	//
	//var s4 = i3[1:2] //无法切片 'i3' (类型 'interface{}')
	//fmt.Println(s4)

	fmt.Println("------------------------")
	var x2 interface{} = 300
	assert(x2)
	var y2 interface{} = "从0到Go语言微服务架构师"
	assert(y2)

	fmt.Println("------------------------")
	getTypeValue(300)
	getTypeValue("从0到Go语言微服务架构师")
	getTypeValue(true)

	fmt.Println("------------------------")
	worker2 := Worker{
		name: "小明",
		book: "从0到Go语言微服务架构师",
		by:   "视频+答疑+群组",
	}
	worker2.learn()
	worker2.rest()
}

func getTypeValue(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("Type:int,Value:%d\n", i.(int))
	case string:
		fmt.Printf("Type:string,Value:%s\n", i.(string))
	default:
		fmt.Printf("Unkown Type\n")
	}
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func ShowType(i interface{}) {
	fmt.Printf("类型:%T,值:%v\n", i, i)
}

// ShowInterface 接口看成type和value的组合，type是接口底层的具体类型，value是具体类型的值
func ShowInterface(s Study) {
	fmt.Printf("接口类型:%T\n,接口值%v\n", s, s)
}

type Life interface {
	Study
	Happy
}

type Study interface {
	learn() //学习
}

type Happy interface {
	rest() //休息
}

type Student struct {
	name string //名字
	book string //课程名称
}

type Worker struct {
	name string //名字
	book string //课程名称
	by   string //学习方式
}

func (s Student) learn() {
	fmt.Printf("%s 在读 %s\n", s.name, s.book)
}

func (s Student) rest() {
	fmt.Printf("%s 放学了，出去玩...\n", s.name)
}

func (w *Worker) learn() {
	fmt.Printf("%s 在读 %s 学习方式 %s\n", w.name, w.book, w.by)
}

func (w *Worker) rest() {
	fmt.Printf("%s 下班了，吃大餐去...\n", w.name)
}
