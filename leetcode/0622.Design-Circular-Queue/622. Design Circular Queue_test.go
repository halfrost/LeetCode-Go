package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem622(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.EnQueue(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.DeQueue()
	fmt.Printf("param_1 = %v obj = %v\n", param2, obj)
	param3 := obj.Front()
	fmt.Printf("param_1 = %v obj = %v\n", param3, obj)
	param4 := obj.Rear()
	fmt.Printf("param_1 = %v obj = %v\n", param4, obj)
	param5 := obj.IsEmpty()
	fmt.Printf("param_1 = %v obj = %v\n", param5, obj)
	param6 := obj.IsFull()
	fmt.Printf("param_1 = %v obj = %v\n", param6, obj)
}
