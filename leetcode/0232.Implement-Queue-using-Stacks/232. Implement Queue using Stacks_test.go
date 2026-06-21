package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem232(t *testing.T) {
	obj := Constructor232()
	fmt.Printf("obj = %v\n", obj)
	obj.Push(2)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(10)
	fmt.Printf("obj = %v\n", obj)
	param2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Peek()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param4)

	// Peek when the Queue is empty so it triggers fromStackToQueue.
	obj2 := Constructor232()
	obj2.Push(5)
	obj2.Push(7)
	peeked := obj2.Peek()
	if peeked != 5 {
		t.Fatalf("Peek() = %v, want %v", peeked, 5)
	}
	fmt.Printf("peeked = %v\n", peeked)
}
