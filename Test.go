package main

import (
	"fmt"
	"net"
	"os"
)

//import (
//	"fmt"
//	"strings"
//	"code.google.com/p/go-tour/wc"
//)
//
//
//func WordCount(s string) map[string]int {
//	s_array := strings.Fields(s)
//	m := make(map[string]int)
//	for i := 0; i < len(s_array); i++ {
//		_, ok := m[s_array[i]]
//		if ok == false {
//			m[s_array[i]] = 1
//		} else {
//			m[s_array[i]]++
//		}
//	}
//	return m
//}
//
//func fibonacci() func() int {
//	var pre, next, sum int
//	pre = 0
//	next = 1
//	count := -1
//	return func() int {
//		count++
//		if count < 2 {
//			return count
//		}
//		sum = pre + next
//		pre = next
//		next = sum
//		return sum
//	}
//}

//func main() {
//
//	var ch1 chan int // ch1是一个正常的channel，不是单向的
//	var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
//	var ch3 <-chan int // ch3是单向channel，只用于读取int数据
//
//	ch4 := make(chan int)
//	ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
//	ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel   左读右写
//
//	f := fibonacci()
//	fmt.Println("***********")
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}

//}

//利用协程实现斐波那契额数列
//func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
//	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
//	go func() {
//		for {
//			x := <-in
//			a <- x
//			b <- x
//			c <- x
//		}
//	}()
//	return a, b, c
//}
//
//func Fib() <-chan int {
//	x := make(chan int, 2)
//	a, b, out := dup3(x)
//	go func() {
//		x <- 0
//		x <- 1
//		<-a  // 等待channel上接收一个值  所以这里也可以写 <-b
//		for {
//			x <- <-a + <-b
//		}
//	}()
//	return out
//
//}
//
//func main() {
//	x := Fib()
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-x)
//	}
//close(ch) 关闭channel
//}

//func main() {
//f,_ := os.Open("");
//defer f.Close()
//r := bufio.NewReader(f)
//s,ok := r.ReadString('\n')
//if ok == nil{
//	fmt.Println(s)
//}
//cmd := exec.Command("/bin/ls", "-l")
//buf, err := cmd.Output()
//if err == nil {
//	str := string(buf)
//	fmt.Println(str)  //这里输出的是一个字节数组
//}
//}

//网络
//定义通道
var ch chan int = make(chan int)

//定义昵称
var nickname string

func reader(conn *net.TCPConn) {          //读取数据
	buff := make([]byte, 128)
	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 8  //这里用来检测channel是否关闭了   向channel中写东西   <-ch从channel中读取东西
			break
		}
		fmt.Println(string(buff[0:j]))
	}
}

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Server is not starting")
		os.Exit(0)
	}

	//为什么不能放到if之前？ err不为nil的话就是painc了 (painc 与 defer 辨析一下！！！)
	defer conn.Close()        //关闭conn

	go reader(conn)

	fmt.Println("请输入昵称")

	fmt.Scanln(&nickname)

	fmt.Println("你的昵称为:", nickname)

	for {
		var msg string
		fmt.Scanln(&msg)
		b := []byte("<" + nickname + ">" + "说:" + msg)
		conn.Write(b)
		//func(c *TCPConn) Write(b[] byte) (n int,err os.Error)  用于发送数据，返回发送的数据长度或者返回错误
		//func (c *TCPConn) Read(b []byte) (n int, err os.Error)用于接收数据，返回接收的长度或者返回错误，是TCPConn的方法
		//select 为非阻塞的
		select {
		case <-ch:
			fmt.Println("Server错误!请重新连接!")
			os.Exit(1)
		default:
			//不加default的话，那么 <-ch 会阻塞for， 下一个输入就没有法进行
		}

	}
}