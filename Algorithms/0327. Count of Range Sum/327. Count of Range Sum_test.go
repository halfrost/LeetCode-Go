package leetcode

import (
	"fmt"
	"testing"
)

type question327 struct {
	para327
	ans327
}

// para 是参数
// one 代表第一个参数
type para327 struct {
	nums  []int
	lower int
	upper int
}

// ans 是答案
// one 代表第一个答案
type ans327 struct {
	one int
}

func Test_Problem327(t *testing.T) {

	qs := []question327{

		question327{
			para327{[]int{-2, 5, -1}, -2, 2},
			ans327{3},
		},

		question327{
			para327{[]int{0, -3, -3, 1, 1, 2}, 3, 5},
			ans327{2},
		},

		question327{
			para327{[]int{-3, 1, 2, -2, 2, -1}, -3, -1},
			ans327{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 327------------------------\n")

	for _, q := range qs {
		_, p := q.ans327, q.para327
		fmt.Printf("【input】:%v       【output】:%v\n", p, countRangeSum(p.nums, p.lower, p.upper))
	}
	fmt.Printf("\n\n\n")
}
