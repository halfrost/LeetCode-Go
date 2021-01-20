package leetcode

import (
	"fmt"
	"testing"
)

type question1672 struct {
	para1672
	ans1672
}

// para 是参数
// one 代表第一个参数
type para1672 struct {
	accounts [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1672 struct {
	one int
}

func Test_Problem1672(t *testing.T) {

	qs := []question1672{

		{
			para1672{[][]int{{1, 2, 3}, {3, 2, 1}}},
			ans1672{6},
		},

		{
			para1672{[][]int{{1, 5}, {7, 3}, {3, 5}}},
			ans1672{10},
		},

		{
			para1672{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
			ans1672{17},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1672------------------------\n")

	for _, q := range qs {
		_, p := q.ans1672, q.para1672
		fmt.Printf("【input】:%v       【output】:%v\n", p, maximumWealth(p.accounts))
	}
	fmt.Printf("\n\n\n")
}
