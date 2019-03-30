package array

import (
	"fmt"
	"testing"
)

//数组声明
func TestArray(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(a[0])

	var b [5]int
	b[0] = 1
	fmt.Println(b[0])

	//用 ... 代替数字
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)
}

//数组截取
func TestSubArray(t *testing.T) {

	var a = [...]int{1, 2, 3, 4, 5}

	fmt.Println(a[1:4])      //[2 3 4]
	fmt.Println(a[1:len(a)]) //[2 3 4 5]
	fmt.Println(a[1:])       //[2 3 4 5]
	fmt.Println(a[:5])       //[1 2 3 4 5]
	fmt.Println(a[:])        //[1 2 3 4 5]
}
