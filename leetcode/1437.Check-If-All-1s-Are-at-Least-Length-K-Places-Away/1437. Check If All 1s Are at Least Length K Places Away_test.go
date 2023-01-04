package leetcode

import (
	"fmt"
	"testing"
)

type question1437 struct {
	para1437
	ans1437
}

// para 是参数
// one 代表第一个参数
type para1437 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1437 struct {
	one bool
}

func Test_Problem1437(t *testing.T) {

	qs := []question1437{

		{
			para1437{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2},
			ans1437{true},
		},

		{
			para1437{[]int{1, 0, 0, 1, 0, 1}, 2},
			ans1437{false},
		},

		{
			para1437{[]int{1, 1, 1, 1, 1}, 0},
			ans1437{true},
		},

		{
			para1437{[]int{0, 1, 0, 1}, 1},
			ans1437{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1437------------------------\n")

	for _, q := range qs {
		_, p := q.ans1437, q.para1437
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", kLengthApart(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
