package main

import (
	"sync"
	"fmt"
	"errors"
)

type myBuffer struct {
	ch           chan interface{} //存放数据的通道
	close        uint32           //标记缓冲器的状态
	closeingLock sync.Mutex       // 为了消除因关闭缓冲器而产生的竞态条件的读写锁
}

func NewBuffer(size uint32) (Buffer, error) {
	if size == 0 {
		errMsg := fmt.Sprintf("illegal size for buffer: %d", size)
		return nil, errors.New(errMsg)
	}
	return &myBuffer{
		ch: make(chan interface{}, size),
	}, nil
}
