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

	obj = Constructor307([]int{-1})
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 0))
	obj.Update(0, 1)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 0))

	obj = Constructor307([]int{7, 2, 7, 2, 0})
	fmt.Printf("obj = %v\n", obj)
	obj.Update(4, 6)
	obj.Update(0, 2)
	obj.Update(0, 9)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(4, 4))
	obj.Update(3, 8)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 4))
	obj.Update(4, 1)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 3))
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 4))
	obj.Update(0, 4)
}
