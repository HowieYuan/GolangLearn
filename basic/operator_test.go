package basic_test

import (
	"fmt"
	"testing"
)

// &^ 运算符
func TestOperator(t *testing.T) {
	// ++a 前缀运算符不支持

	// &^ 按位清零
	a := 7
	a = a &^ 1
	fmt.Println(a)
}

// 循环
func TestFot(t *testing.T) {
	sum := 0

	//传统 for 循环
	for i := 0; i < 10; i++ {
		sum += i
	}

	//while 循环，go 中没有 while 关键字
	n := 0
	for n < 5 {
		sum += 1
		n++
	}

	//foreach 循环数组
	nums := [5]int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	//死循环
	for {
		sum++
	}
}

// if
func TestIf(t *testing.T) {
	num := 1

	//if 中的简短语句
	if v := num; v < 5 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}

func TestSwitch(t *testing.T) {
	s := "aaa"

	//不需要加 break
	//不限制常量和整数，可以是各种类型
	switch s {
	//多个结果选项
	case "aaa", "ccc":
		fmt.Println("yes")
	case "bbb":
		fmt.Println("no")
	default:
		fmt.Println("nil")
	}

	num := 1

	//switch 后面不跟表达式，则效果同 if...else...
	switch {
	case num > 1:
		fmt.Println("> 1")
	case num == 1:
		fmt.Println("= 1")
	case num < 1:
		fmt.Println("< 1")
	default:
		fmt.Println("nil")
	}
}
