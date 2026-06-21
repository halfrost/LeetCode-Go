package leetcode

import (
	"fmt"
	"testing"
)

type question58 struct {
	para58
	ans58
}

// para 是参数
type para58 struct {
	s string
}

// ans 是答案
type ans58 struct {
	ans int
}

func Test_Problem58(t *testing.T) {

	qs := []question58{

		{
			para58{"Hello World"},
			ans58{5},
		},

		{
			para58{"   fly me   to   the moon  "},
			ans58{4},
		},

		{
			para58{"luffy is still joyboy"},
			ans58{6},
		},

		{
			para58{"    "},
			ans58{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 58------------------------\n")

	for _, q := range qs {
		a, p := q.ans58, q.para58
		got := lengthOfLastWord(p.s)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.ans {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.ans, got)
		}
	}
	fmt.Printf("\n\n\n")
}
