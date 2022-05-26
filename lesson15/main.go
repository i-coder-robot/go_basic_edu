package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		通道的声明 chan 就是channel的缩写
		var channel_name chan channel_type
	*/

	//var ch chan string
	//
	//ch = make(chan string)

	//ch := make(chan string)

	/*
		发送数据
		channel_name <- data

		接收数据
		value := <-channel_name
	*/

	//创建一个通道
	//ch := make(chan string)
	////打印 学习课程:
	//fmt.Println("学习课程:")
	////开启协程
	//go PrintChan(ch)
	////从通道中接收数据
	//rec := <-ch
	////打印从通道接收的数据
	//fmt.Println(rec)
	////打印 学习目标:全面掌握Go语言微服务落地，代码级一次性解决微服务和分布式系统。
	//fmt.Println("学习目标:全面掌握Go语言微服务落地，代码级一次性解决微服务和分布式系统。")

	//time.Sleep()

	//关闭通道
	//close(channel_name)
	//value, ok := <-channel_name

	//c := make(chan int, 3)
	//fmt.Println("初始化后：")
	//fmt.Println("cap=", cap(c))
	//fmt.Println("len=", len(c))
	//c <- 1
	//c <- 2
	//fmt.Println("传入2个参数后：")
	//fmt.Println("cap=", cap(c))
	//fmt.Println("len=", len(c))
	//<-c
	//fmt.Println("读取一个数据后：")
	//fmt.Println("cap=", cap(c))
	//fmt.Println("len=", len(c))

	/*
		无缓冲通道
		c:=make(chan int)
		c:=make(chan int,0)

		缓冲通道
		c:=make(chan int,3)
	*/

	//c := make(chan int)
	//go func() {
	//	fmt.Println("send 1")
	//	c <- 1
	//}()
	//
	//go func() {
	//	n := <-c
	//	fmt.Println("receive:", n)
	//}()
	//
	//time.Sleep(1 * time.Second)

	/*
		只读通道 <-chan
		只写通道 chan <-
	*/

	//创建一个双向通道
	//

	//遍历通道  for range
	//var ch2 = make(chan int, 5)
	//go loopPrint(ch2)
	//for v := range ch2 {
	//	fmt.Println(v)
	//}
	//
	//ch3 := make(chan bool, 1)
	//var x int
	//for i := 0; i < 10000; i++ {
	//	go increment(ch3, &x)
	//}
	//time.Sleep(time.Second)
	//fmt.Println("x=", x)

	//ch := make(chan bool)
	//ch <- true //fatal error: all goroutines are asleep - deadlock!

	//ch := make(chan bool)
	//go funcReceiver(ch)
	//ch <- true
	////fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock!
	//time.Sleep(time.Second)

	//ch5 := make(chan bool, 1)
	//ch5 <- true
	//fmt.Println(<-ch5)

	//ch6 := make(chan bool, 1)
	//ch6 <- true
	//ch6 <- false
	//fmt.Println(<-ch6) //fatal error: all goroutines are asleep - deadlock!

	/*
		waitgroup 等待一组任务结束，再执行其他业务逻辑

		Add() 初始值是0，累加子协程的数量。
		Done() 当某个子协程完成后，计数器减去1，通常用defer调用。
		Wait() 阻塞当前协程，直到实例中的计数器归零。
	*/

	//isDone := make(chan bool)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(i)
	//	}
	//	isDone <- true
	//}()
	//
	//<-isDone
	//fmt.Println("main goroutine finished")

	//初始化sync.WaitGroup
	var wg sync.WaitGroup
	//传入子协程的数量
	wg.Add(3)
	//开启子协程1
	go DoTask(1, &wg)
	//开启子协程2
	go DoTask(2, &wg)
	//开启子协程3
	go DoTask(3, &wg)
	//阻塞当前协程，等待所有子协程完成任务
	wg.Wait()
}

func DoTask(taskNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("task: %d: %d\n", taskNum, i)
	}
}

func funcReceiver(c chan bool) {
	fmt.Println(<-c)
}

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1 //不是原子操作，避免多协程进行操作，使用容量为1的通道，达到锁的效果
	<-ch
}

func loopPrint(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

type Sender = chan<- string
type Receiver = <-chan string

func PrintChan(c chan string) {
	c <- "从0到Go语言微服务架构师"
}
