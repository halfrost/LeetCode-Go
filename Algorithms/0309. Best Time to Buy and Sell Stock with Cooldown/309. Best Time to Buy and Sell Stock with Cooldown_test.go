package leetcode

import (
	"fmt"
	"testing"
)

type question309 struct {
	para309
	ans309
}

// para 是参数
// one 代表第一个参数
type para309 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans309 struct {
	one int
}

func Test_Problem309(t *testing.T) {

	qs := []question309{

		question309{
			para309{[]int{}},
			ans309{0},
		},

		question309{
			para309{[]int{2, 1, 4, 5, 2, 9, 7}},
			ans309{10},
		},

		question309{
			para309{[]int{6, 1, 3, 2, 4, 7}},
			ans309{6},
		},

		question309{
			para309{[]int{1, 2, 3, 0, 2}},
			ans309{3},
		},

		question309{
			para309{[]int{7, 1, 5, 3, 6, 4}},
			ans309{5},
		},

		question309{
			para309{[]int{7, 6, 4, 3, 1}},
			ans309{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 309------------------------\n")

	for _, q := range qs {
		_, p := q.ans309, q.para309
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit309(p.one))
	}
	fmt.Printf("\n\n\n")
}
