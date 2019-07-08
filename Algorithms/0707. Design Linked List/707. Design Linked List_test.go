package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem707(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", ML2s(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, ML2s(&obj))
	obj.AddAtHead(38)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtHead(45)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.DeleteAtIndex(2)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtIndex(1, 24)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtTail(36)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtIndex(3, 72)
	obj.AddAtTail(76)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtHead(7)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtHead(36)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtHead(34)
	fmt.Printf("obj = %v\n", ML2s(&obj))
	obj.AddAtTail(91)
	fmt.Printf("obj = %v\n", ML2s(&obj))

	fmt.Printf("\n\n\n")

	obj1 := Constructor()
	fmt.Printf("obj1 = %v\n", ML2s(&obj1))
	param2 := obj1.Get(0)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, ML2s(&obj1))
	obj1.AddAtIndex(1, 2)
	fmt.Printf("obj1 = %v\n", ML2s(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, ML2s(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, ML2s(&obj1))
	obj1.AddAtIndex(0, 1)
	fmt.Printf("obj1 = %v\n", ML2s(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf("param_1 = %v obj1 = %v\n", param2, ML2s(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, ML2s(&obj1))
}

func ML2s(head *MyLinkedList) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
