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

		{
			para76{"ADOBECODEBANC", "ABC"},
			ans76{"BANC"},
		},

		{
			para76{"a", "aa"},
			ans76{""},
		},

		{
			para76{"a", "a"},
			ans76{"a"},
		},

		{
			para76{"", "a"},
			ans76{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 76------------------------\n")

	for _, q := range qs {
		a, p := q.ans76, q.para76
		got := minWindow(p.s, p.p)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
