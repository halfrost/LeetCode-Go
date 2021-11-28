package leetcode

import (
	"fmt"
	"testing"
)

type question301 struct {
	para301
	ans301
}

// s 是参数
type para301 struct {
	s string
}

// ans 是答案
type ans301 struct {
	ans []string
}

func Test_Problem301(t *testing.T) {

	qs := []question301{

		{
			para301{"()())()"},
			ans301{[]string{"(())()", "()()()"}},
		},

		{
			para301{"(a)())()"},
			ans301{[]string{"(a())()", "(a)()()"}},
		},

		{
			para301{")("},
			ans301{[]string{""}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 301------------------------\n")

	for _, q := range qs {
		_, p := q.ans301, q.para301
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeInvalidParentheses(p.s))
	}
	fmt.Printf("\n\n\n")
}
