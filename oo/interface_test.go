package oo_test

import "testing"

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
	p = &Student{}
	t.Log(p.Introduce())
}
