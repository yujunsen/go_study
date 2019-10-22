// Store 和 Load 类函数来提供对数值类型
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown 是通知正在执行的 goroutine 停止工作的标志
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行的时间
	time.Sleep(1 * time.Second)
	// 该停止工作了，安全地设置 shutdown 标志
	fmt.Println("shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 要停止工作了吗？
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutdown %s Down\n", name)
			break
		}
	}
}
