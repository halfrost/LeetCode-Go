package leetcode

import (
	"fmt"
	"testing"
)

type question441 struct {
	para441
	ans441
}

// para 是参数
// one 代表第一个参数
type para441 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans441 struct {
	one int
}

func Test_Problem441(t *testing.T) {

	qs := []question441{

		question441{
			para441{5},
			ans441{2},
		},

		question441{
			para441{8},
			ans441{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 441------------------------\n")

	for _, q := range qs {
		_, p := q.ans441, q.para441
		fmt.Printf("【input】:%v       【output】:%v\n", p, arrangeCoins(p.n))
	}
	fmt.Printf("\n\n\n")
}
