package leetcode

import (
	"fmt"
	"testing"
)

type question132 struct {
	para132
	ans132
}

// para 是参数
// one 代表第一个参数
type para132 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans132 struct {
	one int
}

func Test_Problem132(t *testing.T) {

	qs := []question132{

		{
			para132{"aab"},
			ans132{1},
		},

		{
			para132{""},
			ans132{0},
		},

		{
			para132{"a"},
			ans132{0},
		},

		{
			para132{"aba"},
			ans132{0},
		},

		{
			para132{"aaa"},
			ans132{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 132------------------------\n")

	for _, q := range qs {
		a, p := q.ans132, q.para132
		got := minCut(p.s)
		if got != a.one {
			t.Fatalf("minCut(%q) = %d, want %d", p.s, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
