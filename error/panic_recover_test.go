package error_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

//panic
func TestPanic(t *testing.T) {
	//panic执行后会调用defer
	defer fmt.Println("调用 defer")

	fmt.Println("Start")
	//panic 接受的参数是一个空接口
	panic(errors.New("Something wrong!"))
	//该代码不可达
	fmt.Println("End")
}

//os.Exit
func TestOSExit(t *testing.T) {
	//os.Exit执行后不会调用defer
	defer fmt.Println("调用 defer")

	fmt.Println("Start")
	os.Exit(-1)
	//该代码不可达
	fmt.Println("End")
}

// 类似 try catch 的错误处理
func TestPanic22(t *testing.T) {
	//调用recover()方法
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
