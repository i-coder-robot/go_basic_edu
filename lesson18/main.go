package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		type error interface {
			Error() string
		}
	*/
	//FindFile()

	//a := 100
	//b := -10
	//r, err := area(a, b)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Area=", r)

	//a := 100
	//b := -10
	//r, err := area(a, b)
	//if err != nil {
	//	if err, ok := err.(*areaError); ok {
	//		fmt.Printf("length %d or width %d is less than zero", err.length, err.width)
	//		return
	//	}
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Area=", r)

	//a := 100
	//b := -10
	//r, err := area(a, b)
	//if err != nil {
	//	if err, ok := err.(*areaError); ok {
	//		if err.lengthNegative() {
	//			fmt.Printf("error:长度%d 小于0\n", err.length)
	//		}
	//		if err.widthNegative() {
	//			fmt.Printf("error:宽度%d 小于0\n", err.width)
	//		}
	//		return
	//	}
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Area=", r)

	//panic("panic error")

	//panic recover

	//defer fmt.Println("defer main")
	//myTest()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	outOfArray(20)
	fmt.Println("main...")
}

//func outOfArray(x int) {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	var array [5]int
//	array[x] = 1
//}

func outOfArray(x int) {
	var array [5]int
	array[x] = 1
}

func myTest() {
	defer fmt.Println("defer myTest")
	panic("panic myTest")
}

type areaError struct {
	err    string //错误信息
	length int    //长度
	width  int    //宽度
}

func (e *areaError) Error() string {
	//return fmt.Sprintf("length %d, width %d : %s", e.length, e.width, e.err)
	return e.err
}
func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

//func area(a, b int) (int, error) {
//	if a < 0 || b < 0 {
//		//return 0, errors.New("计算错误，长度或宽度，不能小于0。")
//		//return 0, fmt.Errorf("计算错误，长度%d或宽度%d，不能小于0。", a, b)
//		return 0, &areaError{"length or width is negative.", a, b}
//	}
//	return a * b, nil
//}

func area(length, width int) (int, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += "and width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

type errorString struct {
	s string
}

func New(text string) error {
	return &errorString{s: text}
}

func (e *errorString) Error() string {
	return e.s
}

func FindFile() {
	file, err := os.Open("/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(), "opened successfully")
}
