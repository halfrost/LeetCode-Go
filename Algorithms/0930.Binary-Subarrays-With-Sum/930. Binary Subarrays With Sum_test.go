package leetcode

import (
	"fmt"
	"testing"
)

type question930 struct {
	para930
	ans930
}

// para 是参数
// one 代表第一个参数
type para930 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans930 struct {
	one int
}

func Test_Problem930(t *testing.T) {

	qs := []question930{

		question930{
			para930{[]int{1, 0, 1, 0, 1}, 2},
			ans930{4},
		},

		question930{
			para930{[]int{0, 0, 0, 0, 0}, 0},
			ans930{15},
		},

		question930{
			para930{[]int{1, 0, 1, 1, 1, 1, 0, 1, 0, 1}, 2},
			ans930{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 930------------------------\n")

	for _, q := range qs {
		_, p := q.ans930, q.para930
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSubarraysWithSum(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
