package leetcode

import (
	"fmt"
	"testing"
)

type question28 struct {
	para28
	ans28
}

// para 是参数
// one 代表第一个参数
type para28 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans28 struct {
	one int
}

func Test_Problem28(t *testing.T) {

	qs := []question28{

		question28{
			para28{"abab", "ab"},
			ans28{0},
		},

		question28{
			para28{"hello", "ll"},
			ans28{2},
		},

		question28{
			para28{"", "abc"},
			ans28{0},
		},

		question28{
			para28{"abacbabc", "abc"},
			ans28{5},
		},

		question28{
			para28{"abacbabc", "abcd"},
			ans28{-1},
		},

		question28{
			para28{"abacbabc", ""},
			ans28{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 28------------------------\n")

	for _, q := range qs {
		_, p := q.ans28, q.para28
		fmt.Printf("【input】:%v       【output】:%v\n", p, strStr(p.s, p.p))
	}
	fmt.Printf("\n\n\n")
}
