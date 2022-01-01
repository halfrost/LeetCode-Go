package leetcode

import (
	"fmt"
	"testing"
)

type question1296 struct {
	para1296
	ans1296
}

// para 是参数
type para1296 struct {
	nums []int
	k    int
}

// ans 是答案
type ans1296 struct {
	ans bool
}

func Test_Problem1296(t *testing.T) {

	qs := []question1296{

		{
			para1296{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4},
			ans1296{true},
		},

		{
			para1296{[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3},
			ans1296{true},
		},

		{
			para1296{[]int{1, 2, 3, 4}, 3},
			ans1296{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1296------------------------\n")

	for _, q := range qs {
		_, p := q.ans1296, q.para1296
		fmt.Printf("【input】:%v    【output】:%v\n", p, isPossibleDivide(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
