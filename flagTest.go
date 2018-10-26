package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "请输入端口号：例：1234", "")
	address := flag.String("address", "请输入访问地址：例：127.0.0.1:1234", "")
	//host := flag.String("host", "请输入主机名：例：127.0.0.1", "")
	//name := flag.String("name", "请输入用户名：例：zhangsan", "")
	//password := flag.String("password", "请输入密码：例：123456", "")
	//tls := flag.String("tls","请选择是否加密：0：不加密，1：加密","")

	flag.Parse()
	fmt.Println(*port,*address)
}
