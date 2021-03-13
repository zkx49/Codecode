

import (
	"fmt"
	"sync"
)

/* 为什么高并发且多线程程序数据不加处理就有可能导致数据错乱？
原因：在多核系统中，给定数量的协程会被分配到多个线程里运行，这就有可能出现并行运行！
保证数据正确性的两种方式：
1）互斥锁（sync.Mutex） 2)chan (通道) */
var num int
var mtx sync.Mutex
var wg sync.WaitGroup

func add() {
	mtx.Lock()
	defer mtx.Unlock()
	defer wg.Done()
	num += 1 //给操作的全局变量加锁，想要访问先等锁被释放
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println("num:", num)
}

// 这是保证了数据的正确性
//使用chan
/* var num int

func add(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	ch <- 1
	num += 1
	<-ch
}
func main() {
	ch := make(chan int, 1) //这里为什么是1缓存通道???
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(ch, wg)
	}
	wg.Wait()
	fmt.Println("num:", num)
} */
channel本质上是一个消息队列，主要用于协程之间消息的传递，虽然也可以拿来当互斥锁，但成本更高，这是因为channel的结构决定的。