package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem307(t *testing.T) {
	obj := Constructor307([]int{1, 3, 5})
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 2))
}
