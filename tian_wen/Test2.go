package main

import (
	"fmt"
)

func main() {
	//s := []int{1}
	//one := &s[0]
	//s = append(s, 2)
	//for i := range s {
	//	s[i]++
	//}
	//println(*one, s[0], s[1]) // What's the output?  1  2 3
	//s := make([]int, 512)
	//fmt.Println("Len:", len(s), "Cap:", cap(s))
	//                      //512           //512
	//s = append(s, 0)
	//fmt.Println("Len:", len(s), "Cap:", cap(s))
	                      //513            //1024
	s := make([]int, 1024)
	fmt.Println("Len:", len(s), "Cap:", cap(s))
	//                     1024            1024
	s = append(s, 0)
	fmt.Println("Len:", len(s), "Cap:", cap(s))
	            //        1025             1280
}

//func growslice(t *sliceType, old slice, cap int) slice {
//	// ...
//	for {
//		if old.len < 1024 {
//			newcap += newcap
//		} else {
//			newcap += newcap / 4
//		}i
//		f newcap >= cap {
//			break
//		}
//	}/
//	/ ...
//}
/**
	数组的大小是数组类型的一部分，因此是固定的
	Slice不是数组，它是数组的描述符
	它由一个指向数组（第一个元素），len，cap的指针组成
	当遇到不足的容量时，追加复制基础数组


growlice在追加期间处理切片生长。
传递切片类型，旧切片和所需的新最小容量，
并返回一个至少具有该容量的新切片，以及旧数据复制到它

预分配以获得更好的性能和更少的垃圾
复制发生在[] byte和string之间的转换期间
尽可能避免转换，或者让垃圾丢失性能
明智地选择[] byteor字符串
 */
