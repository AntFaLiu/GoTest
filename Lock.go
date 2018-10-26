package main

import (
	"sync"
	"time"
	"fmt"
)

//var mutex sync.Mutex
//
//func main() {
//	count := 0
//	go func() {
//		mutex.Lock()
//		count++
//	}()
//	time.Sleep(time.Second)
//	fmt.Println("the count is: ", count)
//}

type RWMutex struct {
	w sync.Mutex
	writerSem uint32
	readerSem uint32
	readerCount int32
	readerWait int32
}

func main() {
	var mutex sync.RWMutex
	arr := []int{1, 2, 3}
	go func() {
		fmt.Println("写没锁",arr)
		fmt.Println("Try to lock writing operation.")
		mutex.Lock()   //写锁
		fmt.Println("Writing operation is locked.")

		arr = append(arr, 4)
		fmt.Println("写:",arr)
		fmt.Println("Try to unlock writing operation.")
		mutex.Unlock()
		fmt.Println("Writing operation is unlocked.")
	}()

	go func() {
		fmt.Println("读没锁",arr)
		fmt.Println("Try to lock reading operation.")
		fmt.Println("The len of arr is : ", len(arr))
		mutex.RLock()  //读锁
		fmt.Println("The reading operation is locked.")
		fmt.Println("读:",arr)
		fmt.Println("The len of arr is 读: ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	time.Sleep(time.Second * 2)
	return
}