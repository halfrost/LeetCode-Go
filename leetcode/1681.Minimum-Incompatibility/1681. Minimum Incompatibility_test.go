package leetcode

import (
	"fmt"
	"testing"
)

type question1681 struct {
	para1681
	ans1681
}

// para 是参数
// one 代表第一个参数
type para1681 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1681 struct {
	one int
}

func Test_Problem1681(t *testing.T) {

	qs := []question1681{

		{
			para1681{[]int{1, 2, 1, 4}, 2},
			ans1681{4},
		},

		{
			para1681{[]int{6, 3, 8, 1, 3, 1, 2, 2}, 4},
			ans1681{6},
		},

		{
			para1681{[]int{5, 3, 3, 6, 3, 3}, 3},
			ans1681{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1681------------------------\n")

	for _, q := range qs {
		_, p := q.ans1681, q.para1681
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumIncompatibility(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
