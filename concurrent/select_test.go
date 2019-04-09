package concurrent_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func DoAsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//select
func TestSelect(t *testing.T) {
	select {
	case ret := <-DoAsyncService():
		t.Log(ret)
	//超时机制
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

func TestSelect2(t *testing.T) {
	select {
	case ret := <-DoAsyncService():
		t.Log(ret)
		//超时机制
	case ret := <-AsyncService():
		t.Log(ret)
	default:
		t.Error()
	}

}
