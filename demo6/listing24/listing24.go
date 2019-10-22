// 有缓冲的通道和固定数目的
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //要使用的 goroutine 的数量
	taskLoad         = 10 //要处理工作量
)

var wg sync.WaitGroup

func init() {
	//随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	//有缓冲
	tasks := make(chan string, taskLoad)

	//启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go work(tasks, gr)
	}
	// 增加一组要完成的工作
	for post := 1; post <= numberGoroutines; post++ {
		tasks <- fmt.Sprintf("tasks : %d", post)
	}
	// 当所有工作都处理完时关闭通道
	close(tasks)

	wg.Wait()
}

func work(tasks chan string, worker int) {
	defer wg.Done()
	for {
		//等待任务
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker : %d Shutting Down\n", worker)
			return
		}

		//开始工作
		fmt.Printf("Worker : %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//完成工作
		fmt.Printf("Worker : %d : Completed %s\n", worker, task)

	}
}
