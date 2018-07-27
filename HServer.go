package main

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"os"
	"flag"
)

//var (
//	HHost, HPort string
//)

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}
const HTRUE = "1"
var user map[string]string

func main() {
	HHost := os.Args[1]
	HPort := os.Args[2]
	isTls := os.Args[3]
	http.HandleFunc("/hello", Hello) //url:https://localhost:8080/hello?userName=zhangsan&password=123456
	http.HandleFunc("/index", index)
	if isTls == HTRUE{
		err := http.ListenAndServeTLS(":"+HPort, "/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key", nil) //tls
		if err != nil {
			log.Println(err)
			return
		}
	}else {
		host := flag.String("host", HHost, "listen host")
		port := flag.String("port", HPort, "listen port")
		err := http.ListenAndServe(*host + ":" + *port, nil)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func Hello(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("Hello World"))
	fmt.Println("loginTask is running...")

	//模拟延时
	time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()

	log.Println(req.Body)
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	result := NewBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
		log.Println("登录成功")
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
		log.Println("用户名或密码不正确")
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	w.Write(bytes)
	fmt.Fprint(w, string(bytes))

}

func index(w http.ResponseWriter, req *http.Request) {

	request, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("request :%s\n", request)

	//未知类型的推荐处理方法
	//var f interface{}
	//json.Unmarshal(result, &f)
	//m := f.(map[string]interface{})
	//for k, v := range m {
	//	switch vv := v.(type) {https://httpizza-admin.faas.alpha.elenet.me/route/#!project_name=taco.console_service
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case int:
	//		fmt.Println(k, "is int", vv)
	//	case float64:
	//		fmt.Println(k,"is float64",vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array:")
	//		for i, u := range vv {
	//			fmt.Println(i, u)
	//		}
	//	default:
	//		fmt.Println(k, "is of a type I don't know how to handle")
	//	}
	//}

	//结构已知，解析到结构体

	result := NewBaseJsonBean()
	json.Unmarshal([]byte(request), &user)
	if user["userName"] == "zhangsan" && user["password"] == "123456" {
		result.Code = 100
		result.Message = "登录成功"
		log.Println("登陆成功")

	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
		log.Println("用户名或密码不正确")
	}
	bytes, _ := json.Marshal(result)
	log.Println("bytes:", string(bytes))
	w.Write(bytes)

}
