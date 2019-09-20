package leetcode

import (
	"fmt"
	"testing"
)

type question29 struct {
	para29
	ans29
}

// para 是参数
// one 代表第一个参数
type para29 struct {
	dividend int
	divisor  int
}

// ans 是答案
// one 代表第一个答案
type ans29 struct {
	one int
}

func Test_Problem29(t *testing.T) {

	qs := []question29{

		question29{
			para29{10, 3},
			ans29{3},
		},

		question29{
			para29{7, -3},
			ans29{-2},
		},

		question29{
			para29{-1, 1},
			ans29{-1},
		},

		question29{
			para29{1, -1},
			ans29{-1},
		},

		question29{
			para29{2147483647, 3},
			ans29{715827882},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 29------------------------\n")

	for _, q := range qs {
		_, p := q.ans29, q.para29
		fmt.Printf("【input】:%v    【output】:%v\n", p, divide(p.dividend, p.divisor))
	}
	fmt.Printf("\n\n\n")
}
