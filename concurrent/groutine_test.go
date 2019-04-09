package concurrent_test

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		//调用协程
		go func(i int) {
			//加上这行代码，会出现问题？？
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}(i)
	}
	//time.Sleep(time.Second * 1)
}

//错误写法，这种写法会出现资源竞争
func TestGroutine2(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	//time.Sleep(time.Second * 1)
}
