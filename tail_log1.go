package main

import (
	"os"
	"bytes"
	"fmt"
//	"flag"
	"log"
	"flag"
)

const (
	defaultBufSize = 4096
)

func tail(filename string, n int) (lines []string, err error) {
	f, e := os.Stat(filename)
	if e == nil {
		size := f.Size()
		var fi *os.File
		fi, err = os.Open(filename)
		if err == nil {
			b := make([]byte, defaultBufSize)
			sz := int64(defaultBufSize)
			nn := n
			bTail := bytes.NewBuffer([]byte{})
			istart := size
			flag := true
			for flag {
				if istart < defaultBufSize {
					sz = istart
					istart = 0
					//log.Println("istart:   ",istart)
					//flag = false
				} else {
					istart -= sz
					//log.Println("istart-=sz  :   ",istart)
				}
				_, err = fi.Seek(istart, os.SEEK_SET)   //
				if err == nil {
					mm, e := fi.Read(b)
					log.Println("mm:  ",mm)
					if e == nil && mm > 0 {
						j := mm
						for i := mm - 1; i >= 0; i-- {
							if b[i] == '\n' {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[i+1:j])        //将最后一个字符写入
								j = i
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}
								if (nn == n && bLine.Len() > 0) || nn < n { //skip last "\n"
									lines = append(lines, bLine.String())      //将最后一个字符转换为string放在lines中
									nn --
								}
								if nn == 0 {
									flag = false
									break
								}
							}
						}
						if flag && j > 0 {
							log.Println("j : ",j)
							if istart == 0 {
								//log.Println("000000000")
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[:j])
								if bTail.Len() > 0 {
									log.Println("bTail.Len:  ",bTail.Len)
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}
								lines = append(lines, bLine.String())
								flag = false
							} else {

								bb := make([]byte, bTail.Len())
								copy(bb, bTail.Bytes())
								bTail.Reset()
								bTail.Write(b[:j])
								bTail.Write(bb)
							}
						}

					}

				}

			}
			//func (f *File) Seek(offset int64, whence int) (ret int64, err error)
			//func (f *File) Read(b []byte) (n int, err error) {
		}
		defer fi.Close()

	}
	return
}

func main() {
	file := flag.String("name","/Users/ant_oliu/go/1.8/src/GoTest/shell_test/a.txt","请输入文件名")
	rows := flag.Int("rows",1,"请输入行数")
	flag.Parse()
	//file := "/Users/ant_oliu/go/1.8/src/GoTest/shell_test/a.txt"
	lns, _ := tail(*file, *rows)//查看文件末行
	for _, v := range lns {
		fmt.Println(v)
	}
}
