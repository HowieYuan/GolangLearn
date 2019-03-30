package array_test

import (
	"fmt"
	"testing"
)

//切片
func TestSlice(t *testing.T) {
	slice := []int{2, 3, 5, 7, 11, 13}
	// len cap 两个概念
	fmt.Println(len(slice), cap(slice))

	// append
	slice = append(slice, 100)
	fmt.Println(len(slice), cap(slice))

	// make: len=5  cap=10
	slice = make([]int, 5, 10)
	fmt.Println(len(slice), cap(slice))

	// make: len=5  cap=5
	slice = make([]int, 5)
	fmt.Println(len(slice), cap(slice))
}

//测试 len cap 的增长规律
func TestSliceGrowing(t *testing.T) {
	slice := []int{}
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Println(len(slice), cap(slice))
	}
	/*
	   随着元素的插入，切片底层的连续存储空间（数组）会发生变化，当容量不足时，会指向一个新的数组
	   而数组的长度是 2 ^ n 的规律。

	   这也是为什么 append 函数需要一个等于操作  slice = append(slice, i) 而不是直接 append(slice, i)
	   因为这个过程可能会有一个新的指向

	   因此切片相比数组的好处在于可以有个动态增长，当然这样也会有元素 copy 的代价

	       1 1
	       2 2
	       3 4
	       4 4
	       5 8
	       6 8
	       7 8
	       8 8
	       9 16
	       10 16
	*/
}

//测试切片的底层存储空间以及与 len cap 的关系
func TestSliceUnder(t *testing.T) {
	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := nums[5:7]
	slice2 := nums[2:6]

	// cap 与起始点到底层数组末尾的长度有关
	// len=2  cap=5
	fmt.Println(len(slice1), cap(slice1))
	// len=4  cap=8
	fmt.Println(len(slice2), cap(slice2))

	// 切片修改元素后，底层数组也会被修改，连同有关联的切片
	slice1[0] = 1000
	fmt.Println(nums)
	fmt.Println(slice2)

	//直接 == 的写法不支持
	//if slice1 == slice2 {
	//    fmt.Println("=")
	//}

}
