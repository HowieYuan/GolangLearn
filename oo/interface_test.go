package oo_test

import (
	"fmt"
	"testing"
)

// 接口类型
type People interface {
	Introduce() string
}

// 实现不需要对接口有依赖，这叫做接口的非入侵性
type Student struct {
}

// 结构体方法与接口的方法一直
func (s *Student) Introduce() string {
	return "I am a student"
}

func TestInterface(t *testing.T) {
	var p People
	// 实现接口
	p = new(Student)
	t.Log(p.Introduce())

	// 接口变量
	// People 是接口，对应的是指针类型，所以必须加 &
	p = &Student{}
	t.Log(p.Introduce())
}

//=========================================================

//空接口, 空接口参数可以接受任何数据类型
func emptyInterface(i interface{}) {
	//判断当前参数数据类型
	if value, ok := i.(int); ok {
		fmt.Println(value, ok)
	} else if value, ok := i.(string); ok {
		fmt.Println(value, ok)
	}

	//第二种方式
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	}
}

func TestEmptyInterface(t *testing.T) {
	emptyInterface(10)
	emptyInterface("111")
}
