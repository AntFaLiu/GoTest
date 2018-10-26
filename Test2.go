package main

import "fmt"

//func main()  {
//	for i := 0;i < 5;i++{
//		defer func(n int) {
//			fmt.Printf("%d",n)
//		}(i * 2)
//	}
//}

func main()  {
	for i := 0;i < 5;i++{
		defer func(n int) {
			fmt.Printf("%d",n)
		}(i)
	}
}