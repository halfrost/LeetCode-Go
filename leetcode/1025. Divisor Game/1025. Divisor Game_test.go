package leetcode

import (
	"fmt"
	"testing"
)

type question1025 struct {
	para1025
	ans1025
}

// para 是参数
// one 代表第一个参数
type para1025 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1025 struct {
	one bool
}

func Test_Problem1025(t *testing.T) {

	qs := []question1025{
		question1025{
			para1025{2},
			ans1025{true},
		},

		question1025{
			para1025{3},
			ans1025{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1025------------------------\n")

	for _, q := range qs {
		_, p := q.ans1025, q.para1025
		fmt.Printf("【input】:%v       【output】:%v\n", p, divisorGame(p.one))
	}
	fmt.Printf("\n\n\n")
}
