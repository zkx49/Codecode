package main

import "fmt"

// 两个通道，一个通道写入，等待另一边读出
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		// select 最大的一条限制就是每个case都带有一个IO操作
		// 在select语句中，go语言按顺序从头到尾评估每一个发送和接收的语句
		// 如果其中的任意一语句可以继续执行(即没有被阻塞)，那么就从那些可以执行的语句中任意选择一条来使用。
		// 如果没有任意一条语句可以执行(即所有的通道都被阻塞)，那么有两种可能的情况：
		// 如果给出了default语句，那么就会执行default语句，同时程序的执行会从select语句后的语句中恢复。
		// 如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去。

		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
