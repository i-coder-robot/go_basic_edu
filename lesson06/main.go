package main

import "fmt"

func test01() {
	//声明时没有指定数组的元素值，默认为零值
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
}

func test02() {
	//声明时对数组进行初始化
	var arr1 = [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr1)

	//段声明方式
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr2)
	//部分初始化，未初始化为零值
	arr3 := [5]int{15, 20}
	fmt.Println(arr3)
	//通过指定索引，方便对数组某几个元素赋值
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println(arr4)
	//直接使用...让编译器为我们计算该数组的长度
	arr5 := [...]int{15, 20, 25, 30, 35}
	fmt.Println(arr5)

}

func test03() {
	//特别注意：数组的长度是类型的一部分。[3]int与[5]int不是同一个类型。
	arr1 := [3]int{15, 20, 25}
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 %T\n", arr1)
	fmt.Printf("type of arr2 %T\n", arr2)
}

func test04() {
	//创建多为数组
	arr := [3][2]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"},
	}
	fmt.Println(arr)
}

func arrLength() {
	//len()函数返回数组中元素的个数
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	fmt.Println("数组的长度是:", len(arr))
}

func showArr() {
	//for range ,可以用下划线 _ ，不适用变量
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	for _, value := range arr {
		fmt.Printf("value=%s\n", value)
	}
}

func arrByValue() {
	//Go语言中，数组是值类型，而不是引用类型。
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	copy := arr
	copy[0] = "Golang"
	fmt.Println(arr)
	fmt.Println(copy)
}

func createSlice() {
	//切片[]Type，数组[n]Type
	//方法1 声明整型切片
	var numberList []int
	fmt.Println(numberList)

	//方法2 声明一个空切片
	var numberListEmpty = []int{}
	fmt.Println(numberListEmpty)

	//方法3 make声明方式  make([]Type,size,cap)
	numList := make([]int, 3, 5)
	fmt.Println(numList)
	//指针：是指向第一个切片元素对应的底层数组元素的地址。（切片的第一个元素不一定是数组中第一个元素）
	//长度：切片中的元素个数。
	//容量：从切片的开始位置到底层数据的结尾位置。

	arr := [5]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师", "微服务", "分布式"}
	var s1 = arr[1:4] //数组变量[起始位置:结束位置]（切片中不包含结束位置的元素，也就是取值到结束位置-1）
	fmt.Println(arr)
	fmt.Println(s1)
}

func sliceLenAndCap() {
	s := make([]int, 3, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func outOfSlice() {
	s := make([]int, 3, 5)
	fmt.Println(s[10])
}

func emptySlice() {
	var numberlist []int
	fmt.Println(numberlist == nil) //true

	fmt.Println(len(numberlist) == 0) //true 判断切片是否为空
}

func modifySlice() {
	var arr = [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	s := arr[:] //[0:len(arr)]
	fmt.Println(arr)
	fmt.Println(s)

	s[0] = "Go语言"
	fmt.Println(arr)
	fmt.Println(s)
}

func appendSliceData() {
	s := []string{"Go语言极简一本通"}
	fmt.Println(s)
	fmt.Println(cap(s))
	//追加一个元素
	s = append(s, "Go语言微服务架构核心22讲")
	fmt.Println(s)
	fmt.Println(cap(s))
	//追加2个元素
	s = append(s, "从0到Go语言微服务架构师", "分布式")
	fmt.Println(s)
	fmt.Println(cap(s))
	//追加一个切片
	s = append(s, []string{"微服务", "分布式锁"}...)
	fmt.Println(s)
	fmt.Println(cap(s))
}

func mSlice() {
	numList := [][]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"},
	}
	fmt.Println(numList)
}

func main() {

}
