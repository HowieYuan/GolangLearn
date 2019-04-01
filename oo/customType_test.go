package oo_test

import (
	"fmt"
	"testing"
)

// 将某个函数定义为自定义类型
type IntConv func(num int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		fmt.Println("自定义类型测试函数~")
		return n
	}
}

//自定义 int 类型
type myInt int

func TestMyInt(t *testing.T) {
	var i myInt
	i = 1
	fmt.Println(i)
}
