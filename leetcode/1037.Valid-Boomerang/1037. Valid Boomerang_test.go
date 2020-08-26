package leetcode

import (
	"fmt"
	"testing"
)

type question1037 struct {
	para1037
	ans1037
}

// para 是参数
// one 代表第一个参数
type para1037 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1037 struct {
	one bool
}

func Test_Problem1037(t *testing.T) {

	qs := []question1037{
		{
			para1037{[][]int{{1, 2}, {2, 3}, {3, 2}}},
			ans1037{true},
		},

		{
			para1037{[][]int{{1, 1}, {2, 2}, {3, 3}}},
			ans1037{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1037------------------------\n")

	for _, q := range qs {
		_, p := q.ans1037, q.para1037
		fmt.Printf("【input】:%v       【output】:%v\n", p, isBoomerang(p.one))
	}
	fmt.Printf("\n\n\n")
}
