package leetcode

import (
	"fmt"
	"testing"
)

type question306 struct {
	para306
	ans306
}

// para 是参数
// one 代表第一个参数
type para306 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans306 struct {
	one bool
}

func Test_Problem306(t *testing.T) {

	qs := []question306{

		{
			para306{"112358"},
			ans306{true},
		},

		{
			para306{"199100199"},
			ans306{true},
		},

		{
			para306{"12"},
			ans306{false},
		},

		{
			para306{"1023"},
			ans306{false},
		},

		{
			para306{"100203"},
			ans306{false},
		},

		{
			para306{"1991001990"},
			ans306{false},
		},

		{
			para306{"01234"},
			ans306{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 306------------------------\n")

	for _, q := range qs {
		a, p := q.ans306, q.para306
		got := isAdditiveNumber(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input %v: got %v, want %v", p.one, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
