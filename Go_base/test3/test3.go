package test3

import "fmt"

/*
思考：如何解决协程同步问题？
死锁：两个或两个以上的进程在执行过程中，由于竞争资源或者由于彼此通信而造成的一种阻塞的现象
若无外力的作用，它们都将无法推进下去。
条件：互斥、不可剥夺、循环等待、请求与保持
*/
/* 互斥锁，Mutex结构体 只有两个公开的指针方法，Lock和Unlock
互斥锁的本质是当一个goroutine访问的时候，其余goroutine都不能访问，
这样虽然避免竞争但是也降低了程序的并发性能，从并行变成串行，因而出现读写锁
因为问题出现在写身上，写会改变数据，读不会。
	sync.RWMutex结构体
	1.一组是对写操作的锁定和解锁
	func (*RWMutex)Lock()
	func (*RWMutex)Unlock()
	2.另一组是对读操作的锁定和解锁
	func (*RWMutex)RLock()
	func (*RWMutex)RUnlock()
*/
// 死锁例子：
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println("send")
	go func() {
		<-ch
		fmt.Println("received")
	}()
	fmt.Println("over")
}

/*
var count = 0
var reMutex sync.RWMutex

func read(index int) {
	for {             for循环会一直读写下去
		reMutex.RLock()
		num := count
		fmt.Println(index, "读出", num)
		time.Sleep(time.Millisecond * 300)
		reMutex.RUnlock()
	}
}
func Write(index int) {
	for {
		rand.Seed(time.Now().UnixNano())  设置随机数种子
		data := rand.Intn(500) 获取随机数
		reMutex.Lock()
		count = data
		fmt.Println(index, "写入", data)
		time.Sleep(time.Millisecond * 300)
		reMutex.Unlock()
	}
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

*/
