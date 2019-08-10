package leetcode

import (
	"fmt"
	"testing"
)

type question76 struct {
	para76
	ans76
}

// para 是参数
// one 代表第一个参数
type para76 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans76 struct {
	one string
}

func Test_Problem76(t *testing.T) {

	qs := []question76{

		question76{
			para76{"ADOBECODEBANC", "ABC"},
			ans76{"BANC"},
		},

		question76{
			para76{"a", "aa"},
			ans76{""},
		},

		question76{
			para76{"a", "a"},
			ans76{"a"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 76------------------------\n")

	for _, q := range qs {
		_, p := q.ans76, q.para76
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, minWindow(p.s, p.p))
	}
	fmt.Printf("\n\n\n")
}
