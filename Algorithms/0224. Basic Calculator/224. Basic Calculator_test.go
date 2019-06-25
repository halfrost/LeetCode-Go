package leetcode

import (
	"fmt"
	"testing"
)

type question224 struct {
	para224
	ans224
}

// para 是参数
// one 代表第一个参数
type para224 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans224 struct {
	one int
}

func Test_Problem224(t *testing.T) {

	qs := []question224{

		question224{
			para224{"1 + 1"},
			ans224{2},
		},
		question224{
			para224{" 2-1 + 2 "},
			ans224{3},
		},

		question224{
			para224{"(1+(4+5+2)-3)+(6+8)"},
			ans224{23},
		},

		question224{
			para224{"2-(5-6)"},
			ans224{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 224------------------------\n")

	for _, q := range qs {
		_, p := q.ans224, q.para224
		fmt.Printf("【input】:%v       【output】:%v\n", p, calculate(p.one))
	}
	fmt.Printf("\n\n\n")
}
