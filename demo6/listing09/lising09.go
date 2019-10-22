//竞争
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 捕获 counter 的值
		value := counter
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
		value++
		counter = value
	}
}
