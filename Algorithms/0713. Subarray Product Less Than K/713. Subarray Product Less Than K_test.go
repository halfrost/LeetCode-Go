package leetcode

import (
	"fmt"
	"testing"
)

type question713 struct {
	para713
	ans713
}

// para 是参数
// one 代表第一个参数
type para713 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans713 struct {
	one int
}

func Test_Problem713(t *testing.T) {

	qs := []question713{

		question713{
			para713{[]int{10, 5, 2, 6}, 100},
			ans713{8},
		},

		question713{
			para713{[]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19},
			ans713{18},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 713------------------------\n")

	for _, q := range qs {
		_, p := q.ans713, q.para713
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSubarrayProductLessThanK(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
