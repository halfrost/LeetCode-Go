package leetcode

import (
	"fmt"
	"testing"
)

type question1021 struct {
	para1021
	ans1021
}

// para 是参数
// one 代表第一个参数
type para1021 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1021 struct {
	one string
}

func Test_Problem1021(t *testing.T) {

	qs := []question1021{
		question1021{
			para1021{"(()())(())"},
			ans1021{"()()()"},
		},

		question1021{
			para1021{"(()())(())(()(()))"},
			ans1021{"()()()()(())"},
		},

		question1021{
			para1021{"()()"},
			ans1021{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1021------------------------\n")

	for _, q := range qs {
		_, p := q.ans1021, q.para1021
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeOuterParentheses(p.one))
	}
	fmt.Printf("\n\n\n")
}
