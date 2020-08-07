package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem911(t *testing.T) {
	obj := Constructor911([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Printf("obj.Q[3] = %v\n", obj.Q(3))   // 0
	fmt.Printf("obj.Q[12] = %v\n", obj.Q(12)) // 1
	fmt.Printf("obj.Q[25] = %v\n", obj.Q(25)) // 1
	fmt.Printf("obj.Q[15] = %v\n", obj.Q(15)) // 0
	fmt.Printf("obj.Q[24] = %v\n", obj.Q(24)) // 0
	fmt.Printf("obj.Q[8] = %v\n", obj.Q(8))   // 1

	obj = Constructor911([]int{0, 0, 0, 0, 1}, []int{0, 6, 39, 52, 75})
	fmt.Printf("obj.Q[45] = %v\n", obj.Q(45)) // 0
	fmt.Printf("obj.Q[49] = %v\n", obj.Q(49)) // 0
	fmt.Printf("obj.Q[59] = %v\n", obj.Q(59)) // 0
	fmt.Printf("obj.Q[68] = %v\n", obj.Q(68)) // 0
	fmt.Printf("obj.Q[42] = %v\n", obj.Q(42)) // 0
	fmt.Printf("obj.Q[37] = %v\n", obj.Q(37)) // 0
	fmt.Printf("obj.Q[99] = %v\n", obj.Q(99)) // 0
	fmt.Printf("obj.Q[26] = %v\n", obj.Q(26)) // 0
	fmt.Printf("obj.Q[78] = %v\n", obj.Q(78)) // 0
	fmt.Printf("obj.Q[43] = %v\n", obj.Q(43)) // 0
}
