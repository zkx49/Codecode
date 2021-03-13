package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var count = 0
var reMutex sync.RWMutex

func read(index int) {

	reMutex.RLock()
	num := count
	fmt.Println(index, "读出", num)
	time.Sleep(time.Millisecond * 300)
	reMutex.RUnlock()

}
func Write(index int) {

	rand.Seed(time.Now().UnixNano())
	data := rand.Intn(500)
	reMutex.Lock()
	count = data
	fmt.Println(index, "写入", data)
	time.Sleep(time.Millisecond * 300)
	reMutex.Unlock()

}

// 启动读写协程，
func main() {
	for i := 0; i < 5; i++ {
		go Write(i + 1)
	}
	for j := 0; j < 5; j++ {
		go read(j + 1)
	}
	for {
		runtime.GC() //内存回收
	}
}
