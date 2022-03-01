package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem2166(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf("obj = %v\n", obj)

	obj.Fix(3)
	fmt.Printf("obj = %v\n", obj)
	obj.Fix(1)
	fmt.Printf("obj = %v\n", obj)
	obj.Flip()
	fmt.Printf("obj = %v\n", obj)

	fmt.Printf("all = %v\n", obj.All())
	obj.Unfix(0)
	fmt.Printf("obj = %v\n", obj)
	obj.Flip()
	fmt.Printf("obj = %v\n", obj)

	fmt.Printf("one = %v\n", obj.One())
	obj.Unfix(0)
	fmt.Printf("obj = %v\n", obj)

	fmt.Printf("count = %v\n", obj.Count())
	fmt.Printf("toString = %v\n", obj.ToString())
}
