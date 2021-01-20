package leetcode

import (
	"fmt"
	"testing"
)

type question1690 struct {
	para1690
	ans1690
}

// para 是参数
// one 代表第一个参数
type para1690 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1690 struct {
	one int
}

func Test_Problem1690(t *testing.T) {

	qs := []question1690{

		// {
		// 	para1690{[]int{1, -1, -2, 4, -7, 3}, 2},
		// 	ans1690{7},
		// },

		// {
		// 	para1690{[]int{10, -5, -2, 4, 0, 3}, 3},
		// 	ans1690{17},
		// },

		{
			para1690{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2},
			ans1690{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1690------------------------\n")

	for _, q := range qs {
		_, p := q.ans1690, q.para1690
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxResult(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
