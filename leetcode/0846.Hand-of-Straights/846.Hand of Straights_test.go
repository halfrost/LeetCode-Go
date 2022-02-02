package leetcode

import (
	"fmt"
	"testing"
)

type question846 struct {
	para846
	ans846
}

// para 是参数
type para846 struct {
	hand      []int
	groupSize int
}

// ans 是答案
type ans846 struct {
	ans bool
}

func Test_Problem846(t *testing.T) {

	qs := []question846{

		{
			para846{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3},
			ans846{true},
		},

		{
			para846{[]int{1, 2, 3, 4, 5}, 4},
			ans846{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 846------------------------\n")

	for _, q := range qs {
		_, p := q.ans846, q.para846
		fmt.Printf("【input】:%v    【output】:%v\n", p, isNStraightHand(p.hand, p.groupSize))
	}
	fmt.Printf("\n\n\n")
}
