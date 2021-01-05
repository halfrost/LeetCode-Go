package leetcode

import (
	"fmt"
	"testing"
)

type question830 struct {
	para830
	ans830
}

// para 是参数
// one 代表第一个参数
type para830 struct {
	S string
}

// ans 是答案
// one 代表第一个答案
type ans830 struct {
	one [][]int
}

func Test_Problem830(t *testing.T) {

	qs := []question830{

		{
			para830{"abbxxxxzzy"},
			ans830{[][]int{{3, 6}}},
		},

		{
			para830{"abc"},
			ans830{[][]int{{}}},
		},

		{
			para830{"abcdddeeeeaabbbcd"},
			ans830{[][]int{{3, 5}, {6, 9}, {12, 14}}},
		},

		{
			para830{"aba"},
			ans830{[][]int{{}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 830------------------------\n")

	for _, q := range qs {
		_, p := q.ans830, q.para830
		fmt.Printf("【input】:%v       【output】:%v\n", p, largeGroupPositions(p.S))
	}
	fmt.Printf("\n\n\n")
}
