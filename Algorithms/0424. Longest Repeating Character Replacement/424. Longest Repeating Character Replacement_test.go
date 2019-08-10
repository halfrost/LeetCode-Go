package leetcode

import (
	"fmt"
	"testing"
)

type question424 struct {
	para424
	ans424
}

// para 是参数
// one 代表第一个参数
type para424 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans424 struct {
	one int
}

func Test_Problem424(t *testing.T) {

	qs := []question424{

		question424{
			para424{"AABABBA", 1},
			ans424{4},
		},

		question424{
			para424{"ABBB", 2},
			ans424{4},
		},

		question424{
			para424{"BAAA", 0},
			ans424{3},
		},

		question424{
			para424{"ABCDE", 1},
			ans424{2},
		},

		question424{
			para424{"BAAAB", 2},
			ans424{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 424------------------------\n")

	for _, q := range qs {
		_, p := q.ans424, q.para424
		fmt.Printf("【input】:%v       【output】:%v\n", p, characterReplacement(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
