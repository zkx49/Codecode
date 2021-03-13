

/* 条件变量的作用并不保证在同一时刻仅有一个协程（线程）访问某个共享的数据资源，
而是在对应的共享数据的状态发生变化时，通知阻塞在某个条件上的协程（线程）。
条件变量不是锁，在并发中不能达到同步的目的，因此条件变量总是与锁一块使用。 */

/*
sync.Cond代表了条件变量
func (c *Cond)Wait()
1）阻塞等待条件变量满足
2）释放已掌握的互斥锁相当于cond.L.Unlock(). 注意：两步为一个原子操作
3）当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()
func (c *Cond)Signal()  单发通知
func (c *Cond)Broadcast()  广播通知
*/
func Producer(out chan<- int, cond *sync.Cond) {
	for {
		cond.L.Lock()
		for len(out) == 5 {
			cond.Wait() //阻塞
		}
		num := rand.Intn(500)
		out <- num
		fmt.Println("-------生产数据", num)
		time.Sleep(time.Millisecond * 100)
		cond.L.Unlock()
		cond.Signal()
	}
}
func Consumer(in <-chan int, cond *sync.Cond) {
	for {
		cond.L.Lock()
		if len(in) == 0 {
			cond.Wait() //阻塞
		}
		num := <-in
		fmt.Println("消费数据", num)
		time.Sleep(time.Millisecond * 200)
		cond.L.Unlock()
		cond.Signal() //唤醒
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)
	cond := new(sync.Cond)
	cond.L = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go Producer(ch, cond)
	}
	for j := 0; j < 5; j++ {
		go Consumer(ch, cond)
	}
	for true {
		runtime.GC()
	}
}
