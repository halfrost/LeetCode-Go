package leetcode

import (
	"fmt"
	"testing"
)

type question856 struct {
	para856
	ans856
}

// para 是参数
// one 代表第一个参数
type para856 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans856 struct {
	one int
}

func Test_Problem856(t *testing.T) {

	qs := []question856{
		question856{
			para856{"()"},
			ans856{1},
		},

		question856{
			para856{"(())"},
			ans856{2},
		},

		question856{
			para856{"()()"},
			ans856{2},
		},

		question856{
			para856{"(()(()))"},
			ans856{6},
		},

		question856{
			para856{"()(())"},
			ans856{3},
		},

		question856{
			para856{"((()()))"},
			ans856{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 856------------------------\n")

	for _, q := range qs {
		_, p := q.ans856, q.para856
		fmt.Printf("【input】:%v       【output】:%v\n", p, scoreOfParentheses(p.one))
	}
	fmt.Printf("\n\n\n")
}
