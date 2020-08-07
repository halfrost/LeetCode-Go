package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem303(t *testing.T) {
	obj := Constructor303([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("SumRange(0,2) = %v\n", obj.SumRange(0, 2)) // return 1
	fmt.Printf("SumRange(2,5) = %v\n", obj.SumRange(2, 5)) // return -1
	fmt.Printf("SumRange(0,5) = %v\n", obj.SumRange(0, 5)) // return -3
}
