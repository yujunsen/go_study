//goroutine 调度器行为

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	//wg 用来等待程序完成
	//计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")

	//声明一个匿名函数 并创建goroutine
	go func() {
		//在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		//显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		//在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		//显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

/*
runtime 包的 GOMAXPROCS 函数。这个函数允许程序
更改调度器可以使用的逻辑处理器的数量。如果不想在代码里做这个调用，也可以通过修改和这
个函数名字一样的环境变量的值来更改逻辑处理器的数量。给这个函数传入 1，是通知调度器只
能为该程序使用一个逻辑处理器。

WaitGroup 是一个计数信号量，可以用来记录并维护运行的 goroutine。如果 WaitGroup
的值大于 0， Wait 方法就会阻塞

使用 defer 声明在函数退出时
调用 Done 方法

关键字 defer 会修改函数调用时机，在正在执行的函数返回时才真正调用 defer 声明的函
数。对这里的示例程序来说，我们使用关键字 defer 保证， 每个 goroutine 一旦完成其工作就调
用 Done 方法。

*/
