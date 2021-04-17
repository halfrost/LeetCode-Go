package leetcode

import (
	"fmt"
	"testing"
)

type question1209 struct {
	para1209
	ans1209
}

// para 是参数
// one 代表第一个参数
type para1209 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1209 struct {
	one string
}

func Test_Problem1209(t *testing.T) {

	qs := []question1209{

		// {
		// 	para1209{"abcd", 2},
		// 	ans1209{"abcd"},
		// },

		{
			para1209{"deeedbbcccbdaa", 3},
			ans1209{"aa"},
		},

		{
			para1209{"pbbcggttciiippooaais", 2},
			ans1209{"ps"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1209------------------------\n")

	for _, q := range qs {
		_, p := q.ans1209, q.para1209
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeDuplicates(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
