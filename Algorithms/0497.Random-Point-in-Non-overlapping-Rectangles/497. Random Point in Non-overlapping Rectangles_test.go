package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem497(t *testing.T) {
	w := [][]int{[]int{1, 1, 5, 5}}
	sol := Constructor497(w)
	fmt.Printf("1.Pick = %v\n", sol.Pick())
	fmt.Printf("2.Pick = %v\n", sol.Pick())
	fmt.Printf("3.Pick = %v\n", sol.Pick())
	fmt.Printf("4.Pick = %v\n", sol.Pick())
	fmt.Printf("5.Pick = %v\n", sol.Pick())
	fmt.Printf("6.Pick = %v\n", sol.Pick())
}
