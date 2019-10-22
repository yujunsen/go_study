package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建无缓冲
	court := make(chan int)

	wg.Add(2)

	go palyer("Nadal", court)
	go palyer("Djokovic", court)

	//发球
	court <- 1

	wg.Wait()

}

//模拟选手打网球
func palyer(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//关闭通道 输了
			close(court)
			return
		}
		//显示击球时并加一
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		//将球打给对方
		court <- ball
	}
}
