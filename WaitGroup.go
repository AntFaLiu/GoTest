package main

import (
	"sync"
	"strconv"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	num := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go sum(&num,&wg)
	}

	wg.Wait() // 等待，直到计数为0
	a := strconv.Itoa(num)
	s := a + "dsadasd"
	fmt.Println("s:", s)

	//var wg sync.WaitGroup
	//
	//wg.Add(3) // 因为有两个动作，所以增加2个计数
	//go func() {
	//	fmt.Println("Goroutine 1")
	//	wg.Done() // 操作完成，减少一个计数
	//}()
	//
	//go func() {
	//	fmt.Println("Goroutine 2")
	//	wg.Done() // 操作完成，减少一个计数
	//}()
	//
	//wg.Wait() // 等待，直到计数为0
}

func sum(num *int, wg *sync.WaitGroup) {
	*num++
	wg.Done()
}
