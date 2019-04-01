package functiont_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//一个函数可以有多个返回值
func returnMultiValues() (int, int, int) {
	return rand.Intn(10), rand.Intn(20), rand.Intn(30)
}

func TestMultiValues(t *testing.T) {
	a, b, c := returnMultiValues()
	t.Log(a, b, c)

	// 下划线忽略返回值
	c, _, e := returnMultiValues()
	t.Log(c, e)
}

//=====================================================================

// 函数可以作为参数，也可以作为返回值
// 一个包装函数，类似装饰器模式，包装了函数用时功能
func getTime(function func(i int) int) func(i int) int {
	return func(n int) int {
		start := time.Now()
		result := function(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return result
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := getTime(slowFun)
	t.Log(tsSF(10))
}

//=====================================================================

//可变参数：可以不规定参数的个数，通过 range 函数将参数逐个取出
func Sum(num ...int) int {
	var result int
	for _, val := range num {
		result += val
	}
	return result
}

func TestMultiParameter(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}
