package leetcode

import (
	"fmt"
	"testing"
)

type question1438 struct {
	para1438
	ans1438
}

// para 是参数
// one 代表第一个参数
type para1438 struct {
	nums  []int
	limit int
}

// ans 是答案
// one 代表第一个答案
type ans1438 struct {
	one int
}

func Test_Problem1438(t *testing.T) {

	qs := []question1438{

		{
			para1438{[]int{8, 2, 4, 7}, 4},
			ans1438{2},
		},

		{
			para1438{[]int{10, 1, 2, 4, 7, 2}, 5},
			ans1438{4},
		},

		{
			para1438{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0},
			ans1438{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1438------------------------\n")

	for _, q := range qs {
		_, p := q.ans1438, q.para1438
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestSubarray(p.nums, p.limit))
	}
	fmt.Printf("\n\n\n")
}
