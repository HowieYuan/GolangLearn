package oo_test

import "testing"

type Animal interface {
	say() string
}

type Cat struct {
}

func (cat Cat) say() string {
	return "喵~"
}

type Dog struct {
}

func (dog Dog) say() string {
	return "汪~"
}

func TestPolymorphism(t *testing.T) {
	var a1 Animal = new(Cat)
	var a2 Animal = &Dog{}

	t.Log(a1.say())
	t.Log(a2.say())
}
