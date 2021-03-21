package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem703(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf("Add 7 = %v\n", obj.Add(3))
	fmt.Printf("Add 7 = %v\n", obj.Add(5))
	fmt.Printf("Add 7 = %v\n", obj.Add(10))
	fmt.Printf("Add 7 = %v\n", obj.Add(9))
	fmt.Printf("Add 7 = %v\n", obj.Add(4))
}
