package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem146(t *testing.T) {
	obj := Constructor(2)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.Put(1, 1)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.Put(2, 2)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.Put(3, 3)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 = obj.Get(2)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.Put(4, 4)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 = obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	param1 = obj.Get(3)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	param1 = obj.Get(4)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))

	// Exercise the update-existing-key path in Put (lines 28-33),
	// the Remove-head branch and the Remove-middle branch.
	obj2 := Constructor(3)
	obj2.Put(1, 1)
	obj2.Put(2, 2)
	obj2.Put(3, 3)
	// Update existing key 2 -> triggers Put's "key exists" branch and
	// removes a middle node, then re-adds it at the head.
	obj2.Put(2, 20)
	if v := obj2.Get(2); v != 20 {
		t.Fatalf("expected Get(2)=20, got %v", v)
	}
	fmt.Printf("obj2 = %v\n", MList2Ints(&obj2))
	// Update the current head key to exercise Remove-head branch.
	obj2.Put(2, 200)
	if v := obj2.Get(2); v != 200 {
		t.Fatalf("expected Get(2)=200, got %v", v)
	}
	fmt.Printf("obj2 = %v\n", MList2Ints(&obj2))
}

func MList2Ints(lru *LRUCache) [][]int {
	res := [][]int{}
	head := lru.head
	for head != nil {
		tmp := []int{head.Key, head.Val}
		res = append(res, tmp)
		head = head.Next
	}
	return res
}
