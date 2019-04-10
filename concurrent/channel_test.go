package concurrent_test

import (
	"fmt"
	"sync"
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

//============================================================

//通过消费者生产者模式了解 channel 的关闭机制

func Producer(channel chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i <= 20; i++ {
			channel <- i
		}
		//channel 关闭之后，会发送广播，使得返回 ok 值为 false
		close(channel)

		wg.Done()
	}()
}

func Receiver(channel chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			//根据 ok 值判断通道是否关闭
			if num, ok := <-channel; ok == true {
				fmt.Println(num)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	Producer(channel, &wg)

	//利用通道关闭的广播机制，可以同时设置多个消费者
	wg.Add(1)
	Receiver(channel, &wg)
	wg.Add(1)
	Receiver(channel, &wg)
	wg.Add(1)
	Receiver(channel, &wg)
	wg.Wait()
}

//======================================================

//通过 close 取消 channel，结束所有任务

func Test(t *testing.T) {
	channel := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancelled(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancel")
		}(i, channel)
	}
	cancel_2(channel)
	time.Sleep(time.Second)
}

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

//这种方法只能将消息告诉一个协程
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

//利用 close 可以广播所有协程
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}
