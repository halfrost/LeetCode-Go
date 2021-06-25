package leetcode

import (
	"fmt"
	"testing"
)

type question576 struct {
	para576
	ans576
}

// para 是参数
// one 代表第一个参数
type para576 struct {
	m           int
	n           int
	maxMove     int
	startRow    int
	startColumn int
}

// ans 是答案
// one 代表第一个答案
type ans576 struct {
	one int
}

func Test_Problem576(t *testing.T) {

	qs := []question576{

		{
			para576{2, 2, 2, 0, 0},
			ans576{6},
		},

		{
			para576{1, 3, 3, 0, 1},
			ans576{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 576------------------------\n")

	for _, q := range qs {
		_, p := q.ans576, q.para576
		fmt.Printf("【input】:%v       【output】:%v\n", p, findPaths(p.m, p.n, p.maxMove, p.startRow, p.startColumn))
	}
	fmt.Printf("\n\n\n")
}
