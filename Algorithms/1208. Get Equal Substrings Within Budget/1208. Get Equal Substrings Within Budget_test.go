package leetcode

import (
	"fmt"
	"testing"
)

type question1208 struct {
	para1208
	ans1208
}

// para 是参数
// one 代表第一个参数
type para1208 struct {
	s       string
	t       string
	maxCost int
}

// ans 是答案
// one 代表第一个答案
type ans1208 struct {
	one int
}

func Test_Problem1208(t *testing.T) {

	qs := []question1208{

		question1208{
			para1208{"abcd", "bcdf", 3},
			ans1208{3},
		},

		question1208{
			para1208{"abcd", "cdef", 3},
			ans1208{1},
		},

		question1208{
			para1208{"abcd", "acde", 0},
			ans1208{1},
		},

		question1208{
			para1208{"thjdoffka", "qhrnlntls", 11},
			ans1208{3},
		},

		question1208{
			para1208{"krrgw", "zjxss", 19},
			ans1208{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1208------------------------\n")

	for _, q := range qs {
		_, p := q.ans1208, q.para1208
		fmt.Printf("【input】:%v       【output】:%v\n", p, equalSubstring(p.s, p.t, p.maxCost))
	}
	fmt.Printf("\n\n\n")
}
