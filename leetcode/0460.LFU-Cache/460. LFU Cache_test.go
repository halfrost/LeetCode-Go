package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem460(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(1, 1)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(2, 2)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(3, 3)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(4, 4)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(5, 5)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)

	param1 := obj.Get(4)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(4)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(4)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	param1 = obj.Get(5)
	fmt.Printf("param_1 = %v obj.list = %v obj.map = %v obj.min = %v\n", param1, MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(6, 6)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(7, 7)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
	obj.Put(8, 8)
	fmt.Printf("obj.list = %v obj.map = %v obj.min = %v\n", MLists2Ints(&obj), MList2Ints(&obj), obj.min)
}

func Test_Problem460_Edge(t *testing.T) {
	// capacity == 0: Put returns early, Get always returns -1
	zero := Constructor(0)
	zero.Put(1, 1)
	if got := zero.Get(1); got != -1 {
		t.Fatalf("capacity 0: Get(1) = %v, want -1", got)
	}

	// Get on missing key returns -1
	obj := Constructor(2)
	if got := obj.Get(99); got != -1 {
		t.Fatalf("Get(99) on empty = %v, want -1", got)
	}

	// Put updating an existing key keeps the same key and updates value
	obj.Put(1, 1)
	obj.Put(1, 10) // update existing key 1
	if got := obj.Get(1); got != 10 {
		t.Fatalf("Get(1) after update = %v, want 10", got)
	}

	// Drive the min++ branch in Get: single-key cache, repeated Get bumps frequency
	single := Constructor(1)
	single.Put(1, 1)
	if got := single.Get(1); got != 1 {
		t.Fatalf("single Get(1) = %v, want 1", got)
	}
	if single.min != 2 {
		t.Fatalf("single.min after Get = %v, want 2", single.min)
	}
	if got := single.Get(1); got != 1 {
		t.Fatalf("single Get(1) again = %v, want 1", got)
	}
	if single.min != 3 {
		t.Fatalf("single.min after second Get = %v, want 3", single.min)
	}
}

func MList2Ints(lfu *LFUCache) map[int][][]int {
	res := map[int][][]int{}
	for k, v := range lfu.nodes {
		node := v.Value.(*node)
		arr := [][]int{}
		tmp := []int{node.key, node.value, node.frequency}
		arr = append(arr, tmp)
		res[k] = arr
	}
	return res
}

func MLists2Ints(lfu *LFUCache) map[int][]int {
	res := map[int][]int{}
	for k, v := range lfu.lists {
		tmp := []int{}
		for head := v.Front(); head != nil; head = head.Next() {
			tmp = append(tmp, head.Value.(*node).value)
		}
		res[k] = tmp
	}
	return res
}
