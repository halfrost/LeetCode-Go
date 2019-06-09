package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem232(t *testing.T) {
	obj := Constructor_232()
	fmt.Printf("obj = %v\n", obj)
	obj.Push(2)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(10)
	fmt.Printf("obj = %v\n", obj)
	param_2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param_2)
	param_3 := obj.Peek()
	fmt.Printf("param_3 = %v\n", param_3)
	param_4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param_4)
}
