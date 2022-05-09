package main

import (
	"fmt"
	"math"
	"unsafe"
)

func Integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型是%T,num8的大小 %d，num8的值是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是%T,num16的大小 %d，num16的值是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是%T,num32的大小 %d，num32的值是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是%T,num64的大小 %d，num64的值是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是%T,num的大小 %d，num的值是 %d\n", num, unsafe.Sizeof(num), num)
}

func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint64

	fmt.Printf("num8的类型是%T,num8的大小 %d，num8的值是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是%T,num16的大小 %d，num16的值是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是%T,num32的大小 %d，num32的值是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是%T,num64的大小 %d，num64的值是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是%T,num的大小 %d，num的值是 %d\n", num, unsafe.Sizeof(num), num)
}

func showFloat() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Printf("num1的类型是%T,num1是%g\n", num1, num1)
	fmt.Printf("num2的类型是%T,num1是%g\n", num2, num2)
}

func shoChar() {
	var x byte = 65
	var y uint8 = 65
	fmt.Printf("x=%c\n", x)
	fmt.Printf("y=%c\n", y)
}

func sizeOfChar() {
	//byte是uint8的别名
	var x byte = 65
	fmt.Printf("x=%c\n", x)
	fmt.Printf("x占用%d个字节\n", unsafe.Sizeof(x))

	//rune是int32的别名
	var y rune = 'A'
	fmt.Printf("y=%c\n", y)
	fmt.Printf("y占用%d个字节\n", unsafe.Sizeof(y))
}

func main() {
	//int 32位系统占 32位，64位系统占64
	Integer()
	println("-----------------")
	unsignedInteger()
	println("-----------------")
	showFloat()
	println("-----------------")
	shoChar()
	println("-----------------")
	sizeOfChar()
	println("-----------------")
	var r rune = '牛'
	fmt.Println(r)
	println("-----------------")
	var study string
	study = "《Go语言极简一本通》"
	fmt.Println(study)
	study2 := "《从0到GO语言微服务架构师》"
	fmt.Println(study2)
	var s1 = `
				study = "《Go语言极简一本通》"
				fmt.Println(study)
				study2 := "《从0到GO语言微服务架构师》"
				fmt.Println(study2)
			`
	fmt.Println(s1)

	a := true
	b := false
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("true && false = ", a && b)
	fmt.Println("true || false = ", a || b)

	// 内置的 complex 函数用于构建复数
	var x complex64 = complex(1, 2)
	var y complex128 = complex(3, 4)
	var z complex128 = complex(5, 6)
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)

	// 内建的 real 和 imag 函数分别返回复数的实部和虚部
	fmt.Println("real(x) = ", real(x))
	fmt.Println("imag(x) = ", imag(x))
	fmt.Println("y * z = ", y*z)

	//x := 1 + 2i
	//y := 3 + 4i
	//z := 5 + 6i

	s6 := []string{"Go语言极简一本通"}
	// 追加一个元素
	s6 = append(s6, "从0到Go语言微服务架构师")
	// 追加两个元素
	s6 = append(s6, "Go语言微服务架构师核心22讲", "分布式")
	// 追加一个切片 ... 表示解包不能省略
	s6 = append(s6, []string{"微服务", "分布式锁"}...)
	// 在第一个位置插入一个元素
	s6 = append([]string{"语言概述"}, s6...)
	fmt.Println(s6) // [语言概述 Go语言极简一本通 从0到Go语言微服务架构师 Go语言微服务架构师核心22讲 分布式 微服务 分布式锁]

}
