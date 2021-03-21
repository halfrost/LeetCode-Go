package leetcode

import (
	"fmt"
	"testing"
)

type question1696 struct {
	para1696
	ans1696
}

// para 是参数
// one 代表第一个参数
type para1696 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1696 struct {
	one int
}

func Test_Problem1696(t *testing.T) {

	qs := []question1696{

		{
			para1696{[]int{1, -1, -2, 4, -7, 3}, 2},
			ans1696{7},
		},

		{
			para1696{[]int{10, -5, -2, 4, 0, 3}, 3},
			ans1696{17},
		},

		{
			para1696{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2},
			ans1696{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1696------------------------\n")

	for _, q := range qs {
		_, p := q.ans1696, q.para1696
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxResult(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
