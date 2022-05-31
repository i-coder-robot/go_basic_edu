package main

import (
	"fmt"
	"reflect"
)

//func reflectType(x interface{}) {
//	obj := reflect.TypeOf(x)
//	fmt.Println(obj)
//}

//func reflectType(x interface{}) {
//	typeX := reflect.TypeOf(x)
//	valueX := reflect.ValueOf(x)
//	fmt.Println(typeX)
//	fmt.Println(valueX)
//}

func reflectType(x interface{}) {
	typeX := reflect.TypeOf(x)
	fmt.Println(typeX.Kind())
	fmt.Println(typeX)
}

func main() {
	//var a int64 = 123
	//reflectType(a)
	//var b string = "从0到Go语言微服务架构师"
	//reflectType(b)
	/*
		reflect.Type 表示interface{}的具体类型
		reflect.TypeOf()返回reflect.Type
	*/

	//reflect.Kind()
	//var book book
	//reflectType(book)
	//reflectNumField(book)

	//var book = book{
	//	name:   "《Go语言极简一本通》",
	//	target: "学习Go语言语法，并可以单独完成一个单体服务。",
	//	spend:  8,
	//}
	//reflectNumField(book)

	//反射第一定律- 反射可以将"接口类型变量"转换为"反射类型变量"
	var a interface{} = 3.14
	fmt.Printf("接口变量的类型为 %T,值为 %v\n", a, a)
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("从接口变量到反射对象：Type对象类型为 %T\n", t)
	fmt.Printf("从接口变量到反射对象：Value对象类型为 %T\n", v)

	//反射第二定律- 反射可以将"反射类型对象"转换为"接口对象"

	i := v.Interface()
	fmt.Printf("从反射对象到接口变量:对象类型%T,值为%v\n", i, i)
	//使用类型断言进行转换
	x := v.Interface().(float64)
	fmt.Printf("x的类型为 %T,值为 %v\n", x, x)

	//反射第二定律- 如果要修改"反射类型对象"其值必须是"可写的"
	var c float64 = 3.14
	y := reflect.ValueOf(&c).Elem()
	fmt.Println("是否可写:", y.CanSet())
	y.SetFloat(2.1)
	fmt.Println(y)
}

//func reflectNumField(x interface{}) {
//	if reflect.ValueOf(x).Kind() == reflect.Struct {
//		v := reflect.ValueOf(x)
//		fmt.Println("Number of fields", v.NumField())
//	}
//}

func reflectNumField(x interface{}) {
	if reflect.ValueOf(x).Kind() == reflect.Struct {
		v := reflect.ValueOf(x)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

type book struct {
	name   string
	target string
	spend  int
}
