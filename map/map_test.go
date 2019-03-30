package map_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// string 是 key 的类型
	// int 是 value 的类型
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println(m["one"])
	// map 无法查看 cap
	fmt.Println(m, len(m))

	m["one"] = 1111
	fmt.Println(m["one"])

	// 空 map
	m2 := map[string]int{}
	fmt.Println(m2, len(m2))

	// m2 不等于 nil
	if m2 == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// make
	m3 := make(map[int]int, 10)
	fmt.Println(m3, len(m3))
}

// 测试 map 零值
func TestMapNil(t *testing.T) {
	m := map[int]int{}

	// 对于空的 key，默认值为 0
	fmt.Println(m[0])

	// 可以通过这种方式，获得值和当前 key 是否存在的 bool 值
	value, ok := m[0]
	fmt.Println(value, ok)

	if _, ok := m[0]; ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

// map 遍历
func TestRangeMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
