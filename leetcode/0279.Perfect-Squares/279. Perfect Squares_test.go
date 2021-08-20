package leetcode

import (
	"fmt"
	"testing"
)

type question279 struct {
	para279
	ans279
}

// para 是参数
// one 代表第一个参数
type para279 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans279 struct {
	one int
}

func Test_Problem279(t *testing.T) {

	qs := []question279{

		{
			para279{13},
			ans279{2},
		},
		{
			para279{12},
			ans279{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 279------------------------\n")

	for _, q := range qs {
		_, p := q.ans279, q.para279
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSquares(p.n))
	}
	fmt.Printf("\n\n\n")
}
