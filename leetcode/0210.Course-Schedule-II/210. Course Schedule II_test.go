package leetcode

import (
	"fmt"
	"testing"
)

type question210 struct {
	para210
	ans210
}

// para 是参数
// one 代表第一个参数
type para210 struct {
	one int
	pre [][]int
}

// ans 是答案
// one 代表第一个答案
type ans210 struct {
	one []int
}

func Test_Problem210(t *testing.T) {

	qs := []question210{

		{
			para210{2, [][]int{{1, 0}}},
			ans210{[]int{0, 1}},
		},

		{
			para210{2, [][]int{{1, 0}, {0, 1}}},
			ans210{[]int{0, 1, 2, 3}},
		},

		{
			para210{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			ans210{[]int{0, 1, 2, 3}},
		},

		{
			para210{3, [][]int{{1, 0}, {1, 2}, {0, 1}}},
			ans210{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 210------------------------\n")

	for _, q := range qs {
		_, p := q.ans210, q.para210
		fmt.Printf("【input】:%v       【output】:%v\n", p, findOrder(p.one, p.pre))
	}
	fmt.Printf("\n\n\n")
}
