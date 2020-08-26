package leetcode

import (
	"fmt"
	"testing"
)

type question1051 struct {
	para1051
	ans1051
}

// para 是参数
// one 代表第一个参数
type para1051 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1051 struct {
	one int
}

func Test_Problem1051(t *testing.T) {

	qs := []question1051{

		{
			para1051{[]int{1, 1, 4, 2, 1, 3}},
			ans1051{3},
		},

		{
			para1051{[]int{5, 1, 2, 3, 4}},
			ans1051{5},
		},

		{
			para1051{[]int{1, 2, 3, 4, 5}},
			ans1051{0},
		},

		{
			para1051{[]int{5, 4, 3, 2, 1}},
			ans1051{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1051------------------------\n")

	for _, q := range qs {
		_, p := q.ans1051, q.para1051
		fmt.Printf("【input】:%v       【output】:%v\n", p, heightChecker(p.one))
	}
	fmt.Printf("\n\n\n")
}
