//goroutine 调度器是如何在单个线程上
//切分时间片段
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//1个逻辑处理器
	runtime.GOMAXPROCS(1)
	//计数加2
	wg.Add(2)
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	wg.Wait()
	fmt.Println("Terminating program")
}

//显示5000以内的素数
func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
