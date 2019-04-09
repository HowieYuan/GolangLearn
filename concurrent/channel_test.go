package concurrent_test

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Second * 1)
	fmt.Println("异步任务已完成")
	return "完成任务"
}

func AsyncService() chan string {
	//指明channel只能放string类型
	channel := make(chan string)

	//第二种：buffer channel, 指定 channel 容量，可以避免阻塞
	//channel := make(chan string, 1)

	go func() {
		fmt.Println("异步任务")
		ret := Service()
		//将数据放入 channel
		channel <- ret
		fmt.Println("异步任务结束")
	}()
	return channel
}

func OtherTask() {
	time.Sleep(time.Second * 3)
	fmt.Println("完成其他任务")
}

func TestChannel(t *testing.T) {
	channel := AsyncService()
	OtherTask()
	//从 channel 中等待一个消息，当消息未到达，则会阻塞
	fmt.Println(<-channel)
}
