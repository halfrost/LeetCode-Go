package leetcode

import (
	"fmt"
	"testing"
)

type question560 struct {
	para560
	ans560
}

// para 是参数
// one 代表第一个参数
type para560 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans560 struct {
	one int
}

func Test_Problem560(t *testing.T) {

	qs := []question560{

		{
			para560{[]int{1, 1, 1}, 2},
			ans560{2},
		},

		{
			para560{[]int{1, 2, 3}, 3},
			ans560{2},
		},

		{
			para560{[]int{1}, 0},
			ans560{0},
		},

		{
			para560{[]int{-1, -1, 1}, 0},
			ans560{1},
		},

		{
			para560{[]int{1, -1, 0}, 0},
			ans560{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 560------------------------\n")

	for _, q := range qs {
		_, p := q.ans560, q.para560
		fmt.Printf("【input】:%v       【output】:%v\n", p, subarraySum(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
