package leetcode

import (
	"fmt"
	"testing"
)

type question1674 struct {
	para1674
	ans1674
}

// para 是参数
// one 代表第一个参数
type para1674 struct {
	nums  []int
	limit int
}

// ans 是答案
// one 代表第一个答案
type ans1674 struct {
	one int
}

func Test_Problem1674(t *testing.T) {

	qs := []question1674{

		{
			para1674{[]int{1, 2, 4, 3}, 4},
			ans1674{1},
		},

		{
			para1674{[]int{1, 2, 2, 1}, 2},
			ans1674{2},
		},

		{
			para1674{[]int{1, 2, 1, 2}, 2},
			ans1674{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1674------------------------\n")

	for _, q := range qs {
		_, p := q.ans1674, q.para1674
		fmt.Printf("【input】:%v       【output】:%v\n", p, minMoves(p.nums, p.limit))
	}
	fmt.Printf("\n\n\n")
}
