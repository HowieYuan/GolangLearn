package basic

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
