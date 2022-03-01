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
}
