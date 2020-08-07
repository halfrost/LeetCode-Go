package leetcode

import (
	"fmt"
	"testing"
)

type question1003 struct {
	para1003
	ans1003
}

// para 是参数
// one 代表第一个参数
type para1003 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1003 struct {
	one bool
}

func Test_Problem1003(t *testing.T) {

	qs := []question1003{

		question1003{
			para1003{"aabcbc"},
			ans1003{true},
		},

		question1003{
			para1003{"abcabcababcc"},
			ans1003{true},
		},

		question1003{
			para1003{"abccba"},
			ans1003{false},
		},

		question1003{
			para1003{"cababc"},
			ans1003{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1003------------------------\n")

	for _, q := range qs {
		_, p := q.ans1003, q.para1003
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValid1003(p.s))
	}
	fmt.Printf("\n\n\n")
}
