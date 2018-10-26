package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("/Users/ant_oliu/go/1.8/src/GoTest/shell_test/a.txt") //针对test.log文件
	if err != nil {
		log.Fatalf("Open file fail:%v", err)
	}
	reader := bufio.NewReader(file)
	defer file.Close()
	var i int
	for {

		i++
		log.Println(i)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}
		fmt.Print(string(line))
	}

}
