package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem478(t *testing.T) {
	obj := Constructor(1, 0, 0)
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())

	obj = Constructor(10, 5, -7.5)
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
}
