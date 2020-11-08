package leetcode

import (
	"fmt"
	"testing"
)

type question5562 struct {
	para5562
	ans5562
}

// para 是参数
// one 代表第一个参数
type para5562 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans5562 struct {
	one int
}

func Test_Problem5562(t *testing.T) {

	qs := []question5562{

		{
			para5562{"aab"},
			ans5562{0},
		},

		{
			para5562{"aaabbbcc"},
			ans5562{2},
		},

		{
			para5562{"ceabaacb"},
			ans5562{2},
		},

		{
			para5562{""},
			ans5562{0},
		},

		{
			para5562{"abcabc"},
			ans5562{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5562------------------------\n")

	for _, q := range qs {
		_, p := q.ans5562, q.para5562
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minDeletions(p.s))
	}
	fmt.Printf("\n\n\n")
}
