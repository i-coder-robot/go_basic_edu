package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		select{
			case expression1:
				code
			case expression2:
				code
			default:
				code
		}
	*/
	//SelectDemo01()

	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//ch3 := make(chan string)
	//
	//go task1(ch1)
	//go task2(ch2)
	//go task3(ch3)
	//
	//select {
	//case message1 := <-ch1:
	//	fmt.Println("ch1 received:", message1)
	//case message2 := <-ch2:
	//	fmt.Println("ch2 received:", message2)
	//case message3 := <-ch3:
	//	fmt.Println("ch3 received:", message3)
	//}
	//MockDeadLock()
	//EmptySelect()

	//select-case和switch-case很相似，区别switch是顺序的，select不是顺序的。

	//MockTimeout()

	SelectDemo02()
}

func SelectDemo02() {
	c1 := make(chan string, 2)
	c1 <- "从0到Go语言微服务架构师"
	select {
	case c1 <- "Go语言微服务架构核心22讲":
		fmt.Println("c1 received:", <-c1)
		fmt.Println("c1 received:", <-c1)
	default:
		fmt.Println("channel blocking.")
	}
}

func MockTimeout() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	case <-timeout:
		fmt.Println("Timeout,exit.")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

func EmptySelect() {
	select {}
	//发生死锁 fatal error: all goroutines are asleep - deadlock!
}

func MockDeadLock() {
	//发生死锁 fatal error: all goroutines are asleep - deadlock!
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	}
}

func task1(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Go语言极简一本通"
}

func task2(ch chan string) {
	time.Sleep(7 * time.Second)
	ch <- "Go语言微服务架构核心22讲"
}

func task3(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "从0到Go语言微服务架构师"
}

func SelectDemo01() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "Go语言微服务架构核心22讲"
	ch2 <- "从0到Go语言微服务架构师"
	ch3 <- "Go语言极简一本通"
	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	default:
		fmt.Println("No data received.")
	}
}
