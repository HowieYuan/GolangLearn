package concurrent_example_test

import (
	"runtime"
	"testing"
	"time"
)

//场景：多个协程同时运行，当从其中一个协程获得一个消息时，即可完成任务

func getReply(num int) int {
	//通过 buffer channel 避免协程的阻塞
	channel := make(chan int, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 5)
			channel <- i
		}(i)
	}
	return <-channel
}

func TestReply(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(getReply(10))
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
