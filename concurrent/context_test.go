package concurrent_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled2(ctx context.Context) bool {
	select {
	//接收取消通知
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	//根context，当 context 被取消，他的子 context 也都要被取消
	parentContext := context.Background()
	//子context
	ctx, cancel := context.WithCancel(parentContext)
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled2(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
