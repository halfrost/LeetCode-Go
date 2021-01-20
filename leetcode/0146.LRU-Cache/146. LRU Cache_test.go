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
