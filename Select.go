package main

import "fmt"

//golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {

		case c <- x:
			x, y = y, x+y
		case a := <-quit:
			if a ==2{
				fmt.Println("quit")
			}
			fmt.Println("1")
			return
		}
	}
}
func main() {
	//c := make(chan int)
	num := 0
	quit := make(chan int,20)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 1
	//}()
	//fibonacci(c, quit)
}

func add(num int,c chan int)  {
	num++
}

//接口
type Foo int

func (self Foo) Emit() {
	fmt.Printf("%v", self)
}

type Emitter interface {
	Emit()
}
