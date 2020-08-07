package leetcode

import (
	"fmt"
	"testing"
)

type question921 struct {
	para921
	ans921
}

// para 是参数
// one 代表第一个参数
type para921 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans921 struct {
	one int
}

func Test_Problem921(t *testing.T) {

	qs := []question921{
		question921{
			para921{"())"},
			ans921{1},
		},

		question921{
			para921{"((("},
			ans921{3},
		},

		question921{
			para921{"()"},
			ans921{0},
		},

		question921{
			para921{"()))(("},
			ans921{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 921------------------------\n")

	for _, q := range qs {
		_, p := q.ans921, q.para921
		fmt.Printf("【input】:%v       【output】:%v\n", p, minAddToMakeValid(p.one))
	}
	fmt.Printf("\n\n\n")
}
