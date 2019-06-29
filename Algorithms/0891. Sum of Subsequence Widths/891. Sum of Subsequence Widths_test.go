package leetcode

import (
	"fmt"
	"testing"
)

type question891 struct {
	para891
	ans891
}

// para 是参数
// one 代表第一个参数
type para891 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans891 struct {
	one int
}

func Test_Problem891(t *testing.T) {

	qs := []question891{

		question891{
			para891{[]int{2, 1, 3}},
			ans891{6},
		},

		question891{
			para891{[]int{3, 7, 2, 3}},
			ans891{35},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 891------------------------\n")

	for _, q := range qs {
		_, p := q.ans891, q.para891
		fmt.Printf("【input】:%v       【output】:%v\n", p, sumSubseqWidths(p.one))
	}
	fmt.Printf("\n\n\n")
}
