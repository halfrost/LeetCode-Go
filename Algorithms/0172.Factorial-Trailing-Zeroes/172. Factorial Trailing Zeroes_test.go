package leetcode

import (
	"fmt"
	"testing"
)

type question172 struct {
	para172
	ans172
}

// para 是参数
// one 代表第一个参数
type para172 struct {
	s int
}

// ans 是答案
// one 代表第一个答案
type ans172 struct {
	one int
}

func Test_Problem172(t *testing.T) {

	qs := []question172{

		question172{
			para172{3},
			ans172{0},
		},

		question172{
			para172{5},
			ans172{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 172------------------------\n")

	for _, q := range qs {
		_, p := q.ans172, q.para172
		fmt.Printf("【input】:%v       【output】:%v\n", p, trailingZeroes(p.s))
	}
	fmt.Printf("\n\n\n")
}
