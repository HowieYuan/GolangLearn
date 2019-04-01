package oo_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体
type Person struct {
	id   int
	name string
	age  int
}

// 三种初始化方式
func Test(t *testing.T) {
	person1 := Person{311000, "howie", 21}

	person2 := Person{id: 311123, name: "mike"}

	// 这里返回的是一个指针，即 &Person{}
	person3 := new(Person)
	// 通过指针访问成员变量时不需要 -> 符号
	person3.id = 123456
	person3.name = "james"
	person3.age = 30

	t.Log(person1, person2, person3, *person3)
	t.Log(person1.name)
	t.Logf("person1 is a %T", person1) // oo_test.Person
	t.Logf("person3 is a %T", person3) // *oo_test.Person
}

//==========================================================================

// 结构体的行为函数（该函数同样可以被结构体指针调用）
func (person Person) getDescription() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&person.name))
	return fmt.Sprintf("我叫 "+person.name+", 今年 %d 岁", person.age)
}

// 将该函数依附在对应的指针上（该函数同样可以被结构体调用）
// 推荐使用
func (person *Person) getDescription2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&person.name))
	return fmt.Sprintf("我叫 "+person.name+", 今年 %d 岁", person.age)
}

func TestStructOperations(t *testing.T) {
	p := Person{123, "kobe", 40}

	//通过查看指针可以得知，使用结构体绑定的函数，对象的指针不一致，说明对象有复制的操作
	//而使用结构体指针绑定的函数，对象的指针一致
	//因此一般推荐使用指针绑定
	fmt.Printf("Address is %x\n", unsafe.Pointer(&p.name))

	//调用结构体的行为函数
	s := p.getDescription()
	fmt.Println(s)

	// 指针也同样可以调用该函数
	p2 := &Person{123, "kobe", 40}
	s2 := p2.getDescription()
	fmt.Println(s2)
}
