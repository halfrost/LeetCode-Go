package leetcode

import (
	"fmt"
	"testing"
)

type question50 struct {
	para50
	ans50
}

// para 是参数
// one 代表第一个参数
type para50 struct {
	x float64
	n int
}

// ans 是答案
// one 代表第一个答案
type ans50 struct {
	one float64
}

func Test_Problem50(t *testing.T) {

	qs := []question50{

		question50{
			para50{2.00000, 10},
			ans50{1024.00000},
		},

		question50{
			para50{2.10000, 3},
			ans50{9.26100},
		},

		question50{
			para50{2.00000, -2},
			ans50{0.25000},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 50------------------------\n")

	for _, q := range qs {
		_, p := q.ans50, q.para50
		fmt.Printf("【input】:%v       【output】:%v\n", p, myPow(p.x, p.n))
	}
	fmt.Printf("\n\n\n")
}
