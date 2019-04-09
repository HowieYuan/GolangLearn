package concurrent_test

import (
	"sync"
	"testing"
	"time"
)

//非协程安全
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	t.Log(counter)
}

//加锁
func TestMutex(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			//解锁
			defer func() {
				mut.Unlock()
			}()
			//加锁
			mut.Lock()
			counter++
		}()
	}
	//这个 sleep 的目的是为了等待所有的协程都执行完毕
	time.Sleep(1 * time.Second)
	t.Log(counter)
}

//WaitGroup，类似 Java 中的 join，等待协程的完成
func TestWaitGroup(t *testing.T) {
	var mg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 10000; i++ {
		//计数器加 1
		mg.Add(1)
		go func() {
			//解锁
			defer func() {
				mut.Unlock()
			}()
			//加锁
			mut.Lock()
			counter++
			//计数器减 1
			mg.Done()
		}()
	}
	//等待计数器为 0
	mg.Wait()
	t.Log(counter)
}
