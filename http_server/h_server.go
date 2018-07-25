package main

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"flag"
)


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
	HHost := flag.String("host", "请输入主机名：例：127.0.0.1", "")
	HPort := flag.String("port", "请输入端口号：例：1234", "")
	isTls := flag.String("tls","请选择是否加密：0：不加密，1：加密","")
	flag.Parse()
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/index", index)
	if *isTls == HTRUE{
		err := http.ListenAndServeTLS(":"+*HPort, "/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key", nil) //tls
		if err != nil {
			log.Println(err)
			return
		}
	}else {
		err := http.ListenAndServe(*HHost + ":" + *HPort, nil)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func Hello(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("Hello World"))
	fmt.Println("loginTask is running...")

	//delayed
	time.Sleep(time.Second * 2)

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

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
		log.Println("登录成功")
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
		log.Println("用户名或密码不正确")
	}

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}

func index(w http.ResponseWriter, req *http.Request) {
	request, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("request :%s\n", request)
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
	log.Println("bytes:  ", string(bytes))
	w.Write(bytes)
}
