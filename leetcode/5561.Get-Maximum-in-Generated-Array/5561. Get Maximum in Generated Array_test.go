package leetcode

import (
	"fmt"
	"testing"
)

type question5561 struct {
	para5561
	ans5561
}

// para 是参数
// one 代表第一个参数
type para5561 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans5561 struct {
	one int
}

func Test_Problem5561(t *testing.T) {

	qs := []question5561{

		{
			para5561{7},
			ans5561{3},
		},

		{
			para5561{2},
			ans5561{1},
		},

		{
			para5561{3},
			ans5561{2},
		},

		{
			para5561{0},
			ans5561{0},
		},

		{
			para5561{1},
			ans5561{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5561------------------------\n")

	for _, q := range qs {
		_, p := q.ans5561, q.para5561
		fmt.Printf("【input】:%v      【output】:%v      \n", p, getMaximumGenerated(p.n))
	}
	fmt.Printf("\n\n\n")
}
