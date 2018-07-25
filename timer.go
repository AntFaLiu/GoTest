package main

import (
	"fmt"
	"sync"
	"time"
)

//func main()  {
//	ticker := time.NewTimer(time.Second * 5)
//	select {
//	case <-ticker.C:
//		fmt.Println("1")
//	}
//}
var i int32

func conn() {
	fmt.Println("conn")
	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		//isConn := atomic.LoadInt32(&i)
		//if isConn == 0 {
		//	fmt.Println("000000000000")
		//	//atomic.StoreInt32(&i, 1)
		//}
		fmt.Println("000000000000")
	}

}

var counter int = 0

func Count(lock *sync.Mutex) {
	//lock.Lock() // 上锁
	fmt.Println("333")
	for{
		counter++
		fmt.Println("counter =", counter)
	}

	//lock.Unlock() // 解锁
}

func main() {
	//lock := &sync.Mutex{}
	go conn()

	time.Sleep(5*time.Hour)
	//for i := 0; i < 10; i++ {
	//go Count(lock)
	//for{
	//	fmt.Println("")
	//}
	//runtime.Gosched()
	//for {
	//	lock.Lock() // 上锁
	//	c := counter
	//	lock.Unlock() // 解锁
	//
	//	runtime.Gosched() // 出让时间片
	//
	//	if c >= 10 {
	//		break
	//	}
	//}
}
