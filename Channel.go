package main

import (
	"time"
	"fmt"
)

//var strChan = make(chan string, 3)
//func main()  {
//	var ok bool
//	ch := make(chan int,1)var strChan = make(chan string, 3)

//func main() {
//	synChan1 := make(chan struct{}, 1)
//	synChan2 := make(chan struct{}, 1)
//	go func() {
//		<-synChan1  //一但收到信号就会恢复执行，这时候说明strChan里面已经有了三个值了
//		fmt.Println("Received a sync signal and wait a second...[receiver]")
//		time.Sleep(time.Second)  //睡眠1s是因为strChan的容量是3 当发送第四个值的时候会因为满可而等待，直到接受方取出第一个值
//		//发送方取出第一个值
//		for {
//			if elem, ok := <-strChan; ok {  //ok的作用是判断是否将strChan当中的所用值取出来了
//				fmt.Println("Received:", elem, "[receiver]")
//			} else {
//				break
//			}
//		}
//		fmt.Println("stopped.[receiver]")
//		synChan2 <- struct{}{}
//	}()
//	go func() {
//		for _, elem := range []string{"a", "b", "c","d"} {
//			strChan <- elem
//			fmt.Println("sent:", elem, "[sender]")
//			if elem == "c" {
//				synChan1 <- struct{}{}   //  发送一个信号，这个信号会使接收方恢复执行
//				fmt.Println("sent a sync signal.[sender]")
//			}
//		}
//		fmt.Println("wait 2 second...[sender]")
//		time.Sleep(time.Second * 2)  //睡眠2s是在保证把4个值都接收到
//		close(strChan)
//		synChan2 <- struct{}{}
//	}()
//	<-synChan2
//	<-synChan2   //这两个都是为了不让goroutine 过早的结束
//}
//func main() {
//
//	var ok bool
//	ch := make(chan int, 1)
//	_, ok = interface{}(ch).(<-chan int)
//	fmt.Println("chan int => <-chan int :", ok)
//	_, ok = interface{}(ch).(chan<- int)
//	fmt.Println("chan int => chan int<- :", ok)
//
//	sch := make(chan<- int, 1)
//	_, ok = interface{}(sch).(<-chan int)
//	fmt.Println("chan int => <-chan int :", ok)
//	_, ok = interface{}(sch).(chan int)
//	fmt.Println("chan int => chan int<- :", ok)
//
//	rch := make(<-chan int, 1)
//	_, ok = interface{}(rch).(chan<- int)
//	fmt.Println("chan int => <-chan int :", ok)
//	_, ok = interface{}(rch).(chan int)
//	fmt.Println("chan int => chan int<- :", ok)
//}

//var mapChan = make(chan map[string]int, 1)

//func main() {
//	synChan := make(chan struct{}, 2)
//	go func() {
//		for {
//			if elem, ok := <-mapChan; ok {
//				elem["count"]++
//			} else {
//				break
//			}
//		}
//		fmt.Println("Stopped.[receiver]")
//		synChan <- struct{}{}
//	}()
//
//	go func() {
//		countMap := make(map[string]int)
//		for i := 0; i < 5; i++ {
//			mapChan <- countMap
//			time.Sleep(time.Microsecond)
//			fmt.Printf("the count Map: %v.[sender]\n", countMap)
//		}
//		close(mapChan)
//		synChan <- struct{}{}
//	}()
//	<-synChan
//	<-synChan
//}

//var strChan = make(chan string, 3)
//
//func main() {
//	synChan1 := make(chan struct{}, 1)
//	synChan2 := make(chan struct{}, 2)
//	go receive(strChan, synChan1, synChan2)
//	go send(strChan, synChan1, synChan2)
//	<-synChan2
//	<-synChan2
//
//}
//
//func receive(strChan <-chan string, synChan1 <-chan struct{}, synChan2 chan<- struct{}) {
//	<-synChan1
//	fmt.Println("Received a syn signal and wait a second... [receiver]")
//	time.Sleep(time.Second)
//	for {
//		if elem, ok := <-strChan; ok {
//			fmt.Println("Received: ", elem, "[receiver]")
//		} else {
//			break
//		}
//		fmt.Println("stopped.[receiver]")
//		synChan2 <- struct{}{}
//	}
//}
//
//func send(strChan chan<- string, synChan1 chan<- struct{}, synChan2 chan<- struct{}) {
//	for _, elem := range []string{"a", "b", "c", "d"} {
//		strChan <- elem
//		fmt.Println("sent:", elem, "[sender]")
//		if elem == "c" {
//			synChan1 <- struct{}{}
//			fmt.Println("sent a sync signal.[sender]")
//		}
//	}
//	fmt.Println("wait 2 secounds...[sender]")
//	time.Sleep(time.Second * 2)
//	close(strChan)
//	synChan2 <- struct{}{}
//}

/**
定时器
 */
//func main() {
//	intChan := make(chan int, 1)
//	go func() {
//		for i := 0; i < 5; i++ {
//			time.Sleep(time.Second)
//			intChan <- i   //发送操作有延时
//		}
//		close(intChan)
//	}()
//	timeout := time.Microsecond * 100000
//	var timer *time.Timer
//	for {
//		if timer == nil {
//			timer = time.NewTimer(timeout)
//		} else {
//			fmt.Println("******Reset*****")   //如果定时器不为零，就要先将定时器重置
//			timer.Reset(timeout)
//		}
//		select {
//		case e, ok := <-intChan:
//			if !ok {
//				fmt.Println("end")
//				return
//			}
//			fmt.Printf("Received: %v\n", e)
//		case <-timer.C:
//			fmt.Println("timeout")
//		}
//	}
//}

//func main() {
//	intChan := make(chan int, 1)
//	go func() {
//		time.Sleep(time.Second)
//		intChan <- 1
//	}()
//	select {
//	case e := <-intChan:
//		fmt.Printf("Received: %v\n", e)
//	case <-time.NewTimer(time.Microsecond * 5000000).C:  //由于发送操作有延时，所以如果这里定时器设置的时间太短的话就接收不到
//		fmt.Println("Timeout")
//	}
//}

/**
断续
 */
func main() {
	intChan := make(chan int, 1)
	ticker := time.NewTimer(time.Second)
	go func() {
		for _ = range ticker.C { //每隔一秒随机发送一个【1，3】的数字
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("End.[sender]")
	}()
	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}
	fmt.Println("End.[receiver]")
}
