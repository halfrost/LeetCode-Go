package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem338(t *testing.T) {
	obj := Constructor([]*NestedInteger{})
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj = %v\n", obj.Next())
	fmt.Printf("obj = %v\n", obj.HasNext())
}
