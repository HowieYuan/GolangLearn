package basic_test

import (
	"fmt"
	"testing"
)

// num := 1 该方式不可在函数外部使用

func TestDeclaration(t *testing.T) {
	// 变量声明的几种方式
	var a int
	a = 1
	var b int = 1
	var c = 1
	d := 1

	fmt.Print(a, b, c, d)
}

func TestOperator(t *testing.T) {
	// ++a 前缀运算符不支持

	// &^ 按位清零
	a := 7
	a = a &^ 1
	fmt.Println(a)
}

// 声明自定义类型
type MyInt int

func TestType(t *testing.T) {
	//类型转换
	var a int = 1
	var b int8
	b = int8(a)
	fmt.Println(b)

	c := 1
	d := float64(c)
	fmt.Println(d)

	var e MyInt = 1
	f := int(e)
	fmt.Println(f)
}

func TestConst(t *testing.T) {
	// 常量声明后可以不使用，函数内的变量声明后必须使用，函数外的变量声明后可以不使用
	// 常量无法使用 :=
	const num = 1
}
