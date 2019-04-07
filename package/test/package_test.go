package test

import (
	"fmt"
	"golang_learn/package"
	"testing"
)

func init() {
	fmt.Println("init3")
}

func TestPackage(t *testing.T) {
	packageTest.Test()
}
