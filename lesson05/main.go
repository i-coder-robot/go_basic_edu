package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Integer 有符号整型
func Integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型是 %T,num8的大小是 %d, num8是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T,num16的大小是 %d, num16是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T,num32的大小是 %d, num32是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T,num64的大小是 %d, num64是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T,num的大小是 %d, num是 %d\n", num, unsafe.Sizeof(num), num)
}

// 无符号整型
func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint
	fmt.Printf("num8的类型是 %T, num8的大小 %d, num8是 %d\n",
		num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T, num16的大小 %d, num16是 %d\n",
		num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T, num32的大小 %d, num32是 %d\n",
		num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T, num64的大小 %d, num64是 %d\n",
		num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T, num的大小 %d, num是 %d\n",
		num, unsafe.Sizeof(num), num)
}

func showFloat() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Printf("num1的类型是%T,num1是%g\n", num1, num1)
	fmt.Printf("num2的类型是%T,num2是%g\n", num2, num2)
}

func showChar() {
	var x byte = 65
	var y uint8 = 65

	fmt.Printf("x=%c\n", x)
	fmt.Printf("y=%c\n", y)

}

func sizeOfChar() {
	var x byte = 65
	fmt.Printf("x=%c\n", x)
	fmt.Printf("x占用%d个字节\n", unsafe.Sizeof(x))

	var y rune = 'A'
	fmt.Printf("y=%c\n", y)
	fmt.Printf("y占用%d个字节\n", unsafe.Sizeof(y))
}

func showBool() {
	a := true
	b := false
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("true && false", a && b)
	fmt.Println("true || false", a || b)
	//GO语言中，真-true，假-false,不与数值相等。 (其他语言 true 1 false 0)
}

func showComplex() {
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
}

func main() {
	//一般是用int表示整型的宽度，在32位系统下，就是32位，在64位系统下，就是64位。
	Integer()
	fmt.Println("-----------------")
	unsignedInteger()
	fmt.Println("-----------------")
	showFloat()
	fmt.Println("-----------------")
	showChar()
	fmt.Println("-----------------")
	sizeOfChar()

	var z rune = '牛'
	fmt.Println(z)

	var study string      //定义study字符串型变量
	study = "《Go语言极简一本通》" //赋值
	fmt.Println(study)
	study2 := "《从0到Go语言微服务架构师》" //段声明方式，自动推断类型并初始化
	fmt.Println(study2)

	var s1 string
	s1 = `
			var study string      //定义study字符串型变量
			study = "《Go语言极简一本通》" //赋值
			fmt.Println(study)
		`
	fmt.Println(s1)
	fmt.Println("-----------------")
	showBool()
	fmt.Println("-----------------")
	showComplex()
}
