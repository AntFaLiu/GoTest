package main

//import (
//	"fmt"
//	"time"
//)
//
//func consumer(cname string, ch chan int) {
//	for i := range ch {
//		fmt.Println("consumer--", cname, ":", i)
//	}
//	fmt.Println("ch closed.")
//}
//
//func producer(pname string, ch chan int) {
//	for i := 0; i < 4; i++ {
//
//		fmt.Println("producer--", pname, ":", i)
//		ch <- i
//	}
//}
//
//func main() {
//	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
//	data := make(chan int)
//	go producer("生产者1", data)
//	go producer("生产者2", data)
//	go consumer("消费者1", data)
//	go consumer("消费者2", data)
//
//	time.Sleep(10 * time.Second)
//	close(data)
//	time.Sleep(10 * time.Second)
//}

//import (
//	"fmt"
//)
//
////生产者
//func Producer(ch chan int) {
//	for i := 1; i <= 10; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
////消费者
//func Consumer(id int, ch chan int, done chan bool) {
//	for {
//		value, ok := <-ch //从ch中读取数据
//		if ok {
//			fmt.Printf("id: %d, recv: %d\n", id, value)
//		} else {
//			fmt.Printf("id: %d, closed\n", id)
//			break
//		}
//	}
//	done <- true
//}
//
//func main() {
//	ch := make(chan int, 3)
//	coNum := 2
//	done := make(chan bool, coNum)
//	go Producer(ch)
//	for i := 1; i <= coNum; i++ { //coNum定义的是消费者的数量
//		go Consumer(i, ch, done)
//	}
//	for i := 1; i <= coNum; i++ {
//		<-done    //用来控制协程停止
//	}
//}



import (
	"fmt"
	"math/rand"
	"time"
)

func productor(channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%v", rand.Float64())   //这里会不停的生产
		time.Sleep(time.Second * time.Duration(1))
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string, 5) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	go productor(channel) // 将 productor 函数交给协程处理, 产生的结果传入信道中
	customer(channel) // 主线程从信道中取数据
}
