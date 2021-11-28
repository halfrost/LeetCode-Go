package leetcode

import (
	"fmt"
	"testing"
)

type question488 struct {
	para488
	ans488
}

// para 是参数
type para488 struct {
	board string
	hand  string
}

// ans 是答案
type ans488 struct {
	ans int
}

func Test_Problem488(t *testing.T) {

	qs := []question488{

		{
			para488{"WRRBBW", "RB"},
			ans488{-1},
		},

		{
			para488{"WWRRBBWW", "WRBRW"},
			ans488{2},
		},

		{
			para488{"G", "GGGGG"},
			ans488{2},
		},

		{
			para488{"RBYYBBRRB", "YRBGB"},
			ans488{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 488------------------------\n")

	for _, q := range qs {
		_, p := q.ans488, q.para488
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMinStep(p.board, p.hand))
	}
	fmt.Printf("\n\n\n")
}
