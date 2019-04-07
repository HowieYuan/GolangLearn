package packageTest

import "fmt"

//main 方法执行前，该 package 的所有 init 方法会被优先调用
//一个源文件中 init 方法可以有多个
//一个 package 中的多个 init 方法按照包的导入依赖顺序决定执行顺序
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

//首字母小写无法被外部调用
func test() {
	fmt.Println("test")
}

func Test() {
	fmt.Println("Test")
}
