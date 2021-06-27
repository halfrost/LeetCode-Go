package leetcode

import (
	"fmt"
	"testing"
)

type question135 struct {
	para135
	ans135
}

// para 是参数
// one 代表第一个参数
type para135 struct {
	ratings []int
}

// ans 是答案
// one 代表第一个答案
type ans135 struct {
	one int
}

func Test_Problem135(t *testing.T) {

	qs := []question135{

		{
			para135{[]int{1, 0, 2}},
			ans135{5},
		},

		{
			para135{[]int{1, 2, 2}},
			ans135{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 135------------------------\n")

	for _, q := range qs {
		_, p := q.ans135, q.para135
		fmt.Printf("【input】:%v       【output】:%v\n", p, candy(p.ratings))
	}
	fmt.Printf("\n\n\n")
}
