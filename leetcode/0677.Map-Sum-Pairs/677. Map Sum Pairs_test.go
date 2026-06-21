package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem677(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", obj)
	obj.Insert("apple", 3)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
	obj.Insert("banana", 10)
	fmt.Printf("obj = %v\n", obj)
	if got := obj.Sum("ap"); got != 5 {
		t.Fatalf("Sum(\"ap\") = %d, want 5", got)
	}
	if got := obj.Sum("ba"); got != 10 {
		t.Fatalf("Sum(\"ba\") = %d, want 10", got)
	}
	fmt.Printf("obj.sum = %v\n", obj.Sum("ba"))
}
