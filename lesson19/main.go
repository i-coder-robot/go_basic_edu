package main

import "fmt"

func main() {
	num := new(int)
	fmt.Println(*num)

	//a := make([]int, 2, 10)
	//b := make(map[string]int)
	//c := make(chan int, 10)

	/*
		new 为所有类型分配内存，并且初始化零值，返回指针。
		make 只能为 slice,map,chan分配内存，并且初始化，返回的是类型。
	*/
}
