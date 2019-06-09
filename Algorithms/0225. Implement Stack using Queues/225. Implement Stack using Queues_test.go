package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem225(t *testing.T) {
	obj := Constructor_225()
	fmt.Printf("obj = %v\n", obj)
	param_5 := obj.Empty()
	fmt.Printf("param_5 = %v\n", param_5)
	obj.Push(2)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(10)
	fmt.Printf("obj = %v\n", obj)
	param_2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param_2)
	param_3 := obj.Top()
	fmt.Printf("param_3 = %v\n", param_3)
	param_4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param_4)
}
