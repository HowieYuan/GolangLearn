package concurrent_example_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//等待所有协程完成既可以使用 WaitGroup，也可以通过循环的方式取出每个协程 channel 的消息

func getAllReply(num int) int {
	//通过 buffer channel 避免协程的阻塞
	channel := make(chan int, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 5)
			channel <- i
		}(i)
	}
	result := 0
	//循环获得 channel 中的消息
	for i := 0; i < num; i++ {
		num := <-channel
		result += num
		fmt.Println("get number:", num)
	}
	return result
}

func TestAllReply(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(getAllReply(10))
	//这里就不再需要 sleep 了
	t.Log("After:", runtime.NumGoroutine())
}
