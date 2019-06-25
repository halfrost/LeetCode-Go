package leetcode

import (
	"fmt"
	"testing"
)

type question52 struct {
	para52
	ans52
}

// para 是参数
// one 代表第一个参数
type para52 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans52 struct {
	one int
}

func Test_Problem52(t *testing.T) {

	qs := []question52{

		question52{
			para52{1},
			ans52{1},
		},

		question52{
			para52{2},
			ans52{0},
		},
		question52{
			para52{3},
			ans52{0},
		},

		question52{
			para52{4},
			ans52{2},
		},

		question52{
			para52{5},
			ans52{10},
		},

		question52{
			para52{6},
			ans52{4},
		},

		question52{
			para52{7},
			ans52{40},
		},

		question52{
			para52{8},
			ans52{92},
		},

		question52{
			para52{9},
			ans52{352},
		},

		question52{
			para52{10},
			ans52{724},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 52------------------------\n")
	for _, q := range qs {
		_, p := q.ans52, q.para52
		fmt.Printf("【input】:%v       【output】:%v\n", p, totalNQueens(p.one))
	}
	fmt.Printf("\n\n\n")
}
