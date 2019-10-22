//创建一个 goroutine 池并完成工作
package main

import (
	"log"
	"sync"
	"time"

	"github.com/goinaction/mycode/demo7/patterns/work"
)

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//"github.com/goinaction/mycode/demo7/patterns/runner"
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代 names 切片
		for _, name := range names {
			// 创建一个 namePrinter 并提供
			// 指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行。当 Run 返回时
				// 我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
