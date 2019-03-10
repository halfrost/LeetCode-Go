package leetcode

import (
	"fmt"
	"testing"
)

type question143 struct {
	para143
	ans143
}

// para 是参数
// one 代表第一个参数
type para143 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans143 struct {
	one []int
}

func Test_Problem143(t *testing.T) {

	qs := []question143{

		question143{
			para143{[]int{1, 2, 3, 4, 5}},
			ans143{[]int{1, 5, 2, 4, 3}},
		},
		question143{
			para143{[]int{1, 2, 3, 4}},
			ans143{[]int{1, 4, 2, 3}},
		},

		question143{
			para143{[]int{1}},
			ans143{[]int{1}},
		},

		question143{
			para143{[]int{}},
			ans143{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 143------------------------\n")

	for _, q := range qs {
		_, p := q.ans143, q.para143
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(reorderList(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
