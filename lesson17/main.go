package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// BankV3 银行
type BankV3 struct {
	balance int          //余额
	rwMutex sync.RWMutex //读写锁，适用场景：读多写少（微博，B站）
	/*
		读锁 RLock() RUnlock()
		写锁 Lock() Unlock()
	*/
}

// Deposit 存款
func (b *BankV3) Deposit(amount int) {
	b.rwMutex.Lock()         //加锁
	defer b.rwMutex.Unlock() //解锁
	b.balance += amount
}

// Balance 查询余额
func (b *BankV3) Balance() (balance int) {
	b.rwMutex.RLock()
	balance = b.balance
	b.rwMutex.RUnlock()
	return
}

// BankV2 银行
type BankV2 struct {
	balance int        //余额
	m       sync.Mutex //互斥锁
}

// Deposit 存款
func (b *BankV2) Deposit(amount int) {
	b.m.Lock()         //加锁
	defer b.m.Unlock() //解锁
	b.balance += amount
}

// Balance 查询余额
func (b *BankV2) Balance() int {
	return b.balance
}

// Bank 银行
type Bank struct {
	balance int //余额
}

// Deposit 存款
func (b *Bank) Deposit(amount int) {
	b.balance += amount
}

// Balance 查询余额
func (b *Bank) Balance() int {
	return b.balance
}

func main() {
	//b := &Bank{}
	//b.Deposit(1000)
	//b.Deposit(1000)
	//b.Deposit(1000)
	//
	//fmt.Println(b.Balance())

	//var wg sync.WaitGroup
	//b := &Bank{}
	//n := 1000
	//wg.Add(n)
	//for i := 1; i <= n; i++ {
	//	go func() {
	//		b.Deposit(1000)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(b.Balance()) //968000 950000 956000

	//var wg sync.WaitGroup
	//b := &BankV2{}
	//n := 1000
	//wg.Add(n)
	//for i := 1; i <= n; i++ {
	//	go func() {
	//		b.Deposit(1000)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(b.Balance()) //1000000

	//var wg sync.WaitGroup
	//b := &BankV3{}
	//n := 1000
	//wg.Add(n)
	//for i := 1; i <= n; i++ {
	//	go func() {
	//		b.Deposit(1000)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(b.Balance()) //1000000

	//sync.Cond
	s1 := []string{"张三"}
	s2 := []string{"赵四"}
	s3 := []string{"刘能"}

	var m sync.Mutex
	cond := sync.NewCond(&m) //创建一个带锁的条件变量，Locker通常是Mutex或者RWMutex
	/*
		Broadcast 唤醒所有因等待条件变量c阻塞的goroutine
		Signal 唤醒一个因等待条件变量c阻塞的goroutine
		Wait 等待c.L解锁并挂起goroutine，在稍后恢复执行后，Wait返回锁定c.L,只有当被Broadcast或Signal唤醒，wait才返回。
	*/

	go listen("Go语言极简一本通", s1, cond)
	go listen("Go语言微服务架构核心22讲", s2, cond)
	go listen("从0到Go语言微服务架构师", s3, cond)

	go broadcast("秒杀开始:", cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func listen(name string, s []string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name, "报名:", s)
	c.L.Unlock()
}

func broadcast(event string, c *sync.Cond) {
	time.Sleep(time.Second)
	c.L.Lock()
	fmt.Println(event)
	c.Broadcast()
	c.L.Unlock()
}
