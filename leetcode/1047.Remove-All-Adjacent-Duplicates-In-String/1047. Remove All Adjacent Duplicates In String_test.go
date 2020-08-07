package leetcode

import (
	"fmt"
	"testing"
)

type question1047 struct {
	para1047
	ans1047
}

// para 是参数
// one 代表第一个参数
type para1047 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1047 struct {
	one string
}

func Test_Problem1047(t *testing.T) {

	qs := []question1047{

		question1047{
			para1047{"abbaca"},
			ans1047{"ca"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1047------------------------\n")

	for _, q := range qs {
		_, p := q.ans1047, q.para1047
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeDuplicates1047(p.s))
	}
	fmt.Printf("\n\n\n")
}
