package leetcode

import (
	"fmt"
	"testing"
)

type question1663 struct {
	para1663
	ans1663
}

// para 是参数
// one 代表第一个参数
type para1663 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1663 struct {
	one string
}

func Test_Problem1663(t *testing.T) {

	qs := []question1663{

		{
			para1663{3, 27},
			ans1663{"aay"},
		},

		{
			para1663{5, 73},
			ans1663{"aaszz"},
		},

		{
			para1663{24, 552},
			ans1663{"aaszz"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1663------------------------\n")

	for _, q := range qs {
		_, p := q.ans1663, q.para1663
		fmt.Printf("【input】:%v      【output】:%v      \n", p, getSmallestString(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
