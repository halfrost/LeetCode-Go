package template

import (
	"fmt"
	"testing"
)

func Test_BIT(t *testing.T) {
	nums, bit := []int{1, 2, 3, 4, 5, 6, 7, 8}, BinaryIndexedTree{}
	bit.Init(8)
	fmt.Printf("nums = %v bit = %v\n", nums, bit.tree) // [0 1 3 3 10 5 11 7 36]
}
