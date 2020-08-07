package leetcode

import (
	"fmt"
	"testing"
)

type question1202 struct {
	para1202
	ans1202
}

// para 是参数
// one 代表第一个参数
type para1202 struct {
	s     string
	pairs [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1202 struct {
	one string
}

func Test_Problem1202(t *testing.T) {

	qs := []question1202{

		question1202{
			para1202{"dcab", [][]int{[]int{0, 3}, []int{1, 2}}},
			ans1202{"bacd"},
		},

		question1202{
			para1202{"dcab", [][]int{[]int{0, 3}, []int{1, 2}, []int{0, 2}}},
			ans1202{"abcd"},
		},

		question1202{
			para1202{"cba", [][]int{[]int{0, 1}, []int{1, 2}}},
			ans1202{"abc"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1202------------------------\n")

	for _, q := range qs {
		_, p := q.ans1202, q.para1202
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestStringWithSwaps(p.s, p.pairs))
	}
	fmt.Printf("\n\n\n")
}
