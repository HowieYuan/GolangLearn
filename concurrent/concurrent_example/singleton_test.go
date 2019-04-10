package concurrent_example_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

//golang 实现单例模式
//利用 Once 实现 “只执行一次” 机制

type Singleton struct {
	//data string
}

var instance *Singleton

// 只执行一次动作的对象
var once sync.Once

func GetInstance() *Singleton {
	//执行唯一一次的动作
	once.Do(func() {
		instance = new(Singleton)
		fmt.Println("create!")
	})
	return instance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			instance := GetInstance()
			//打印出来的地址都相同
			fmt.Println(unsafe.Pointer(instance))
			wg.Done()
		}()
	}
	wg.Wait()
}
