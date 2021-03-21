package leetcode

import (
	"fmt"
	"testing"
)

type question785 struct {
	para785
	ans785
}

// para 是参数
// one 代表第一个参数
type para785 struct {
	graph [][]int
}

// ans 是答案
// one 代表第一个答案
type ans785 struct {
	one bool
}

func Test_Problem785(t *testing.T) {

	qs := []question785{

		{
			para785{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}},
			ans785{true},
		},

		{
			para785{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}},
			ans785{false},
		},

		{
			para785{[][]int{{1, 2, 3}, {0, 2, 3}, {0, 1, 3}, {0, 1, 2}}},
			ans785{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 785------------------------\n")

	for _, q := range qs {
		_, p := q.ans785, q.para785
		fmt.Printf("【input】:%v       【output】:%v\n", p, isBipartite(p.graph))
	}
	fmt.Printf("\n\n\n")
}
