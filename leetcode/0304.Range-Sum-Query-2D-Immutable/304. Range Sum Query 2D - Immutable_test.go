package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem304(t *testing.T) {
	obj := Constructor(
		[][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
	)
	fmt.Printf("obj = %v\n", obj.SumRegion(2, 1, 4, 3))
	fmt.Printf("obj = %v\n", obj.SumRegion(1, 1, 2, 2))
	fmt.Printf("obj = %v\n", obj.SumRegion(1, 2, 2, 4))

	empty := Constructor([][]int{})
	if empty.cumsum != nil {
		t.Fatalf("expected nil cumsum for empty matrix, got %v", empty.cumsum)
	}
}
