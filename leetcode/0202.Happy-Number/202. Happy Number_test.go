package leetcode

import (
	"fmt"
	"testing"
)

type question202 struct {
	para202
	ans202
}

// para 是参数
// one 代表第一个参数
type para202 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans202 struct {
	one bool
}

func Test_Problem202(t *testing.T) {

	qs := []question202{

		question202{
			para202{202},
			ans202{false},
		},

		question202{
			para202{19},
			ans202{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 202------------------------\n")

	for _, q := range qs {
		_, p := q.ans202, q.para202
		fmt.Printf("【input】:%v       【output】:%v\n", p, isHappy(p.one))
	}
	fmt.Printf("\n\n\n")
}
