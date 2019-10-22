//4 个 goroutine接力赛
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//无缓冲通道
	baton := make(chan int)

	//最后一位跑步者加一
	wg.Add(1)

	//第一位跑步者持有接力棒
	go Runner(baton)

	//开始
	baton <- 1

	wg.Wait()

}

func Runner(baton chan int) {
	var newRunner int
	//等待接力棒
	runner := <-baton
	//绕着跑道
	fmt.Printf("Runner %d Running With Baton\n", runner)
	//创建下一位
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	//比赛是否接结束
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	//接力棒交给下一位
	fmt.Printf("Runner %d Exchange With Runner %d \n", runner, newRunner)

	baton <- newRunner
}
