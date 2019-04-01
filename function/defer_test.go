package functiont_test

import (
	"fmt"
	"testing"
)

// defer 推迟调用，类似 Java 的 finally
// 可以用于释放资源，释放锁等
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("清理资源！")
	}()

	t.Log("开始任务")
}

func TestDefer2(t *testing.T) {
	defer fmt.Println("清理资源!")

	t.Log("开始任务")
}
