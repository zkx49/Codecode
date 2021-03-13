package main

import (
	"fmt"
	"time"
)

// 每个case语句必须是一个IO操作
func main() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select { //因为select的机制，所有会选择一条case来执行
			case v := <-c: //这里已经阻塞
				fmt.Println(v)
			case <-time.After(5 * time.Second): //这里给它传入
				fmt.Println("timeout")
				quit <- true //传入数据
				return
			}
		}
	}()
	<-quit //读出

}
