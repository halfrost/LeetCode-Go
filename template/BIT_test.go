package template

import (
	"fmt"
	"testing"
)

func Test_BIT(t *testing.T) {
	nums, bit := []int{1, 2, 3, 4, 5, 6, 7, 8}, BinaryIndexedTree{}
	bit.Init(nums, 8)
	fmt.Printf("%v\n", bit.tree) // [0 1 3 3 10 5 11 7 36]
}
