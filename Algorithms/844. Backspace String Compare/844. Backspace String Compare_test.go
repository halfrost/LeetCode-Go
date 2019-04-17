package leetcode

import (
	"fmt"
	"testing"
)

type question844 struct {
	para844
	ans844
}

// para 是参数
// one 代表第一个参数
type para844 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans844 struct {
	one bool
}

func Test_Problem844(t *testing.T) {

	qs := []question844{

		question844{
			para844{"ab#c", "ad#c"},
			ans844{true},
		},

		question844{
			para844{"ab##", "c#d#"},
			ans844{true},
		},

		question844{
			para844{"a##c", "#a#c"},
			ans844{true},
		},

		question844{
			para844{"a#c", "b"},
			ans844{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 844------------------------\n")

	for _, q := range qs {
		_, p := q.ans844, q.para844
		fmt.Printf("【input】:%v       【output】:%v\n", p, backspaceCompare(p.s, p.t))
	}
	fmt.Printf("\n\n\n")
}
