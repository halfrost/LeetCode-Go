package leetcode

import (
	"fmt"
	"testing"
)

type question1319 struct {
	para1319
	ans1319
}

// para 是参数
// one 代表第一个参数
type para1319 struct {
	n           int
	connections [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1319 struct {
	one int
}

func Test_Problem1319(t *testing.T) {

	qs := []question1319{

		{
			para1319{4, [][]int{{0, 1}, {0, 2}, {1, 2}}},
			ans1319{1},
		},

		{
			para1319{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}},
			ans1319{2},
		},

		{
			para1319{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}},
			ans1319{-1},
		},

		{
			para1319{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}},
			ans1319{0},
		},

		{
			para1319{12, [][]int{{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2}}},
			ans1319{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1319------------------------\n")

	for _, q := range qs {
		_, p := q.ans1319, q.para1319
		fmt.Printf("【input】:%v       【output】:%v\n", p, makeConnected(p.n, p.connections))
	}
	fmt.Printf("\n\n\n")
}
