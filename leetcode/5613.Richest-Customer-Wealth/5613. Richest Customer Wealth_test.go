package leetcode

import (
	"fmt"
	"testing"
)

type question491 struct {
	para491
	ans491
}

// para 是参数
// one 代表第一个参数
type para491 struct {
	accounts [][]int
}

// ans 是答案
// one 代表第一个答案
type ans491 struct {
	one int
}

func Test_Problem491(t *testing.T) {

	qs := []question491{

		{
			para491{[][]int{{1, 2, 3}, {3, 2, 1}}},
			ans491{6},
		},

		{
			para491{[][]int{{1, 5}, {7, 3}, {3, 5}}},
			ans491{10},
		},

		{
			para491{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
			ans491{17},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 491------------------------\n")

	for _, q := range qs {
		_, p := q.ans491, q.para491
		fmt.Printf("【input】:%v       【output】:%v\n", p, maximumWealth(p.accounts))
	}
	fmt.Printf("\n\n\n")
}
