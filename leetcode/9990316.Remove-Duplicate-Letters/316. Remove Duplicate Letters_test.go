package leetcode

import (
	"fmt"
	"testing"
)

type question316 struct {
	para316
	ans316
}

// para 是参数
// one 代表第一个参数
type para316 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans316 struct {
	one string
}

func Test_Problem316(t *testing.T) {

	qs := []question316{

		{
			para316{"bcabc"},
			ans316{"abc"},
		},
		{
			para316{"cbacdcbc"},
			ans316{"acdb"},
		},
		{
			para316{""},
			ans316{""},
		},
		{
			para316{"abcd"},
			ans316{"abcd"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 316------------------------\n")

	for _, q := range qs {
		a, p := q.ans316, q.para316
		got := removeDuplicateLetters(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input %q: expected %q, got %q", p.one, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
