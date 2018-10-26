package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Post("https://api-push.meizu.com/garcia/api/server/push/statistics/dailyPushStatics",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"time"
//	"io/ioutil"
//)
//
//func main() {
//	//生成client 参数为默认
//	client := &http.Client{}
//
//	date := time.Now()
//	fmt.Println(date)
//	//生成要访问的url
//	url := "https://restapi.getui.com/v1/{5a1282a74415cd5fefe97113}/query_app_user/{2018-04-10 16:10:40}"
//
//	//提交请求
//	reqest, err := http.NewRequest("get", url, nil)
//
//	if err != nil {
//		panic(err)
//	}
//
//	//处理返回结果
//	response, _ := client.Do(reqest)
//
//	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
//	stdout := os.Stdout
//	_, err = io.Copy(stdout, response.Body)
//
//	//返回的状态码
//	status := response.StatusCode
//	Body,err := ioutil.ReadAll(response.Body)
//	if(err!=nil){
//		return
//	}
//
//	fmt.Println("Body: " , string(Body))
//
//	fmt.Println(status)
//
//}
