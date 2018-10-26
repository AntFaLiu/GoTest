package main

import (
	"fmt"
	//"time"
)

//import "fmt"
//
//func main() {
//	arr := [10]int{}
//	for i := 0; i < 10; i++ {
//		fmt.Print("Result of ", i, ":")
//		go func() {
//			arr[i] = i
//			fmt.Println("go: ",arr[i])
//		}()
//	}
//	fmt.Println("Done")
//}

//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//	var cnt int = 0         // 全局计数器
//	mylock := &sync.Mutex{} // 互斥锁
//	arr := [10]int{}
//	for i := 0; i < 10; i++ {
//		fmt.Print("Result of ", i, ":")
//		go func() {
//			arr[i] = i + i*i
//			fmt.Println(arr[i])
//			mylock.Lock() // 写之前枷锁
//			cnt++
//			mylock.Unlock() // 写之后释放锁
//		}()
//	}
//	for {
//		mylock.Lock() // 读之前枷锁
//		temp := cnt
//		mylock.Unlock()   // 读之后释放锁
//		runtime.Gosched() // 协程切换
//		if temp >= 10 {
//			break
//		}
//	}
//	fmt.Println("Done")
//}

//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//	var cnt int = 0         // 全局计数器
//	mylock := &sync.Mutex{} // 互斥锁
//	arr := [11]int{}
//	for i := 0; i < 10; i++ {
//		go func() {
//			arr[i] = i + i*i
//			mylock.Lock() // 写之前枷锁
//			cnt++
//			mylock.Unlock() // 写之后释放锁
//		}()
//	}
//	for {
//		mylock.Lock() // 读之前枷锁
//		temp := cnt
//		mylock.Unlock()   // 读之后释放锁
//		runtime.Gosched() // 协程切换
//		if temp >= 10 {
//			break
//		}
//	}
//	for i := 0; i < 11; i++ {
//		fmt.Println("Result of ", i, ":", arr[i])
//	}
//	fmt.Println("Done")
//}

//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//	var cnt int = 0         // 全局计数器
//	mylock := &sync.Mutex{} // 互斥锁
//	arr := [11]int{}
//	for i := 0; i < 10; i++ {
//		go func(i int) { // 这里的i是形参
//			arr[i] = i + i*i
//			mylock.Lock() // 写之前枷锁
//			cnt++
//			mylock.Unlock() // 写之后释放锁
//		}(i) // 这里的i是实参
//	}
//	for {
//		mylock.Lock() // 读之前枷锁
//		temp := cnt
//		mylock.Unlock()   // 读之后释放锁
//		runtime.Gosched() // 协程切换
//		if temp >= 10 {
//			break
//		}
//	}
//	for i := 0; i < 11; i++ {
//		fmt.Println("Result of ", i, ":", arr[i])
//	}
//	fmt.Println("Done")
//}

//import "fmt"
//
//func main() {
//	chs := make([]chan int, 10) // 申请一个10维的channel数组
//	arr := [11]int{}
//	for i := 0; i < 10; i++ {
//		chs[i] = make(chan int) // 对每个channel初始化
//		go func(ch chan int, i int) {
//			arr[i] = i + i*i
//			ch <- 10 // 写channel，应该在函数体的最后一行，许书（P94，代码清单4-4）上有bug
//		}(chs[i], i)
//	}
//	for _, ch := range chs {
//		<-ch // 读channel           //保证channel都收到值
//	}
//	for i := 0; i < 11; i++ {
//		fmt.Println("Result of ", i, ":", arr[i])
//	}
//	fmt.Println("Done")
//}


//func main() {
//	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
//	match := make(chan string, 1) // 为一个未匹配的发送操作提供空间
//	wg := new(sync.WaitGroup)
//	wg.Add(len(people))
//	for _, name := range people {
//		go Seek(name, match, wg)
//	}
//	wg.Wait()
//	select {
//	case name := <-match:
//		fmt.Printf("No one received %s’s message.\n", name)
//	default:
//		// 没有待处理的发送操作
//	}
//}
//
//// 函数Seek 发送一个name到match管道或从match管道接收一个peer，结束时通知wait group
//func Seek(name string, match chan string, wg *sync.WaitGroup) {
//	select {
//	case peer := <-match:
//		fmt.Printf("%s sent a message to %s.\n", peer, name)
//	case match <- name:
//		// 等待某个goroutine接收我的消息
//	}
//	wg.Done()
//}

//
//func race() {
//	wait := make(chan struct{})
//	n := 0
//	go func() {
//		// 译注：注意下面这一行
//		n++ // 一次访问: 读, 递增, 写
//		close(wait)
//	}()
//	// 译注：注意下面这一行
//	n++ // 另一次冲突的访问
//	<-wait
//	fmt.Println(n) // 输出：未指定
//}
//func sharingIsCaring() {
//	ch := make(chan int)
//	go func() {
//		n := 0 // 仅为一个goroutine可见的局部变量.
//		n++
//		ch <- n // 数据从一个goroutine离开...
//	}()
//	n := <-ch   // ...然后安全到达另一个goroutine.
//	n++
//	fmt.Println(n) // 输出: 2
//}
// AtomicInt是一个并发数据结构，持有一个整数值
// 该数据结构的零值为0
//type AtomicInt struct {
//	mu sync.Mutex // 锁，一次仅能被一个goroutine持有。
//	n  int
//}
//
//// Add方法作为一个原子操作将n加到AtomicInt
//func (a *AtomicInt) Add(n int) {
//	a.mu.Lock() // 等待锁释放，然后持有它
//	a.n += n
//	a.mu.Unlock() // 释放锁
//}
//
//// Value方法返回a的值
//func (a *AtomicInt) Value() int {
//	a.mu.Lock()
//	n := a.n
//	a.mu.Unlock()
//	return n
//}
//
//func lockItUp() {
//	wait := make(chan struct{})
//	var n AtomicInt
//	go func() {
//		n.Add(1) // 一个访问
//		close(wait)
//	}()
//	n.Add(1) // 另一个并发访问
//	<-wait
//	fmt.Println(n.Value()) // 输出: 2
//}
//
//func main()  {
//	//race()
//	//sharingIsCaring()
//	lockItUp()
//}

//func race() {
//	var wg sync.WaitGroup
//	wg.Add(5)
//	// 译注：注意下面这行代码中的i++
//	for i := 0; i < 5; i++ {
//		go func() {
//			// 注意下一行代码会输出什么？为什么？
//			fmt.Print(i) // 6个goroutine共享变量i
//			wg.Done()
//		}()
//	}
//	wg.Wait() // 等待所有（5个）goroutine运行结束
//	fmt.Println()
//}
//func correct() {
//	var wg sync.WaitGroup
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		go func(n int) { // 使用局部变量
//			fmt.Print(n)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println()
//}
//func alsoCorrect() {
//	var wg sync.WaitGroup
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		n := i // 为每个闭包创建一个独有的变量
//		go func() {
//			fmt.Print(n)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println()
//}
//
//func main()  {
//	race()
//	//correct()
//	//alsoCorrect()
//}

//type Vector []float64
//func Convolve(u, v Vector) (w Vector) {
//	n := len(u) + len(v) - 1
//	w = make(Vector, n)
//
//	// 将 w 切分成花费 ~100μs-1ms 用于计算的工作单元
//	size := max(1, 1<<20/n)
//
//	wg := new(sync.WaitGroup)
//	wg.Add(1 + (n-1)/size)
//	for i := 0; i < n && i >= 0; i += size { // 整型溢出后 i < 0
//		j := i + size
//		if j > n || j < 0 { // 整型溢出后 j < 0
//			j = n
//		}
//
//		// 这些goroutine共享内存，但是只读
//		go func(i, j int) {
//			for k := i; k < j; k++ {
//				w[k] = mul(u, v, k)
//			}
//			wg.Done()
//		}(i, j)
//	}
//	wg.Wait()
//	return
//}
// 函数mul 返回 Σ u[i]*v[j], i + j = k.
//func mul(u, v Vector, k int) (res float64) {
//	n := min(k+1, len(u))
//	j := min(k, len(v)-1)
//	for i := k - j; i < n; i, j = i+1, j-1 {
//		res += u[i] * v[j]
//	}
//	return
//}

//func main()  {
//	names := []string{"Tom","Jack","Roose","Liasa"}
//	for _,name := range names {
//		go func(){
//			fmt.Println("ddd")
//			fmt.Println("Hello,%s!\n",name)
//		}()
//		time.Sleep(time.Microsecond)
//	}
//	//time.Sleep(time.Microsecond)
//}

func main()  {
	names := []string{"Tom","Jack","Roose","Liasa"}
	for _,name := range names {
		go func(who string){

			fmt.Printf("Hello,%s!\n",who)
		}(name)

	}
	//time.Sleep(time.Microsecond)
}