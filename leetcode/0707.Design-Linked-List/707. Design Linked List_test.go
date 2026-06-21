package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem707(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.AddAtHead(38)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(45)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.DeleteAtIndex(2)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtIndex(1, 24)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtTail(36)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtIndex(3, 72)
	obj.AddAtTail(76)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(7)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(36)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(34)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtTail(91)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))

	fmt.Printf("\n\n\n")

	obj1 := Constructor()
	fmt.Printf("obj1 = %v\n", MList2Ints(&obj1))
	param2 := obj1.Get(0)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, MList2Ints(&obj1))
	obj1.AddAtIndex(1, 2)
	fmt.Printf("obj1 = %v\n", MList2Ints(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, MList2Ints(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, MList2Ints(&obj1))
	obj1.AddAtIndex(0, 1)
	fmt.Printf("obj1 = %v\n", MList2Ints(&obj1))
	param2 = obj1.Get(0)
	fmt.Printf("param_1 = %v obj1 = %v\n", param2, MList2Ints(&obj1))
	param2 = obj1.Get(1)
	fmt.Printf("param_2 = %v obj1 = %v\n", param2, MList2Ints(&obj1))

	fmt.Printf("\n\n\n")

	// obj2 exercises remaining branches: AddAtTail on empty list,
	// DeleteAtIndex(0) on a multi-element list, and DeleteAtIndex on a
	// middle element via the else branch.
	obj2 := Constructor()
	obj2.AddAtTail(1) // head == nil branch in AddAtTail
	if got := MList2Ints(&obj2); len(got) != 1 || got[0] != 1 {
		t.Fatalf("AddAtTail on empty list = %v, want [1]", got)
	}
	obj2.AddAtTail(2)
	obj2.AddAtTail(3)
	fmt.Printf("obj2 = %v\n", MList2Ints(&obj2))
	obj2.DeleteAtIndex(1) // else branch, curr.Next != nil, new curr.Next != nil
	if got := MList2Ints(&obj2); len(got) != 2 || got[0] != 1 || got[1] != 3 {
		t.Fatalf("DeleteAtIndex(1) = %v, want [1 3]", got)
	}
	obj2.DeleteAtIndex(0) // index == 0 branch, head.Next != nil
	if got := MList2Ints(&obj2); len(got) != 1 || got[0] != 3 {
		t.Fatalf("DeleteAtIndex(0) = %v, want [3]", got)
	}
	obj2.DeleteAtIndex(0) // index == 0 branch, head.Next == nil
	if got := MList2Ints(&obj2); len(got) != 0 {
		t.Fatalf("DeleteAtIndex(0) = %v, want []", got)
	}
	fmt.Printf("obj2 = %v\n", MList2Ints(&obj2))
}

func MList2Ints(head *MyLinkedList) []int {
	res := []int{}
	cur := head.head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
