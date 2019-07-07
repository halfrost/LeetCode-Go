package leetcode

import (
	"fmt"
	"testing"
)

type question136 struct {
	para136
	ans136
}

// para 是参数
// one 代表第一个参数
type para136 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans136 struct {
	one int
}

func Test_Problem136(t *testing.T) {

	qs := []question136{

		question136{
			para136{[]int{2, 2, 1}},
			ans136{1},
		},

		question136{
			para136{[]int{4, 1, 2, 1, 2}},
			ans136{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 136------------------------\n")

	for _, q := range qs {
		_, p := q.ans136, q.para136
		fmt.Printf("【input】:%v       【output】:%v\n", p, singleNumber(p.s))
	}
	fmt.Printf("\n\n\n")
}
