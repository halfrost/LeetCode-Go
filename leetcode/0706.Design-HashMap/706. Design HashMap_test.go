package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem706(t *testing.T) {
	obj := Constructor706()
	obj.Put(7, 10)
	fmt.Printf("Get 7 = %v\n", obj.Get(7))
	obj.Put(7, 20)
	fmt.Printf("Contains 7 = %v\n", obj.Get(7))
	param1 := obj.Get(100)
	fmt.Printf("param1 = %v\n", param1)
	obj.Remove(100007)
	param1 = obj.Get(7)
	fmt.Printf("param1 = %v\n", param1)
	obj.Remove(7)
	param1 = obj.Get(7)
	fmt.Printf("param1 = %v\n", param1)

	// Exercise hash collisions to build a chain (7, 10007, 20007 all hash to 7).
	obj2 := Constructor706()
	obj2.Put(7, 1)      // head
	obj2.Put(10007, 2)  // chained node, next == nil branch
	obj2.Put(20007, 3)  // chained node, next != nil recursion
	obj2.Put(10007, 22) // update existing chained key (N.key == key branch in chain)
	if got := obj2.Get(7); got != 1 {
		t.Fatalf("Get(7) = %v, want 1", got)
	}
	if got := obj2.Get(10007); got != 22 {
		t.Fatalf("Get(10007) = %v, want 22", got)
	}
	if got := obj2.Get(20007); got != 3 {
		t.Fatalf("Get(20007) = %v, want 3", got)
	}
	// Get a key that hashes to 7 but is absent: traverse chain to the end -> -1.
	if got := obj2.Get(30007); got != -1 {
		t.Fatalf("Get(30007) = %v, want -1", got)
	}
	// Remove a chained (non-head) node, exercising recursive Remove.
	obj2.Remove(10007)
	if got := obj2.Get(10007); got != -1 {
		t.Fatalf("Get(10007) after remove = %v, want -1", got)
	}
	if got := obj2.Get(20007); got != 3 {
		t.Fatalf("Get(20007) after remove = %v, want 3", got)
	}
	// Remove the head node of a chain.
	obj2.Remove(7)
	if got := obj2.Get(7); got != -1 {
		t.Fatalf("Get(7) after remove = %v, want -1", got)
	}
	if got := obj2.Get(20007); got != 3 {
		t.Fatalf("Get(20007) after head remove = %v, want 3", got)
	}
	// Remove a non-existent chained key (recursion where key not found in chain).
	obj2.Remove(40007)
	// Remove a key whose bucket is empty (HashNode == nil branch in MyHashMap.Remove).
	obj2.Remove(9999)
	fmt.Printf("Get 20007 = %v\n", obj2.Get(20007))
}
