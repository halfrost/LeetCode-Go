package leetcode

import (
	"fmt"
	"testing"
)

type question1752 struct {
	para1752
	ans1752
}

// para 是参数
// one 代表第一个参数
type para1752 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1752 struct {
	one bool
}

func Test_Problem1752(t *testing.T) {

	qs := []question1752{

		{
			para1752{[]int{3, 4, 5, 1, 2}},
			ans1752{true},
		},

		{
			para1752{[]int{2, 1, 3, 4}},
			ans1752{false},
		},

		{
			para1752{[]int{1, 2, 3}},
			ans1752{true},
		},

		{
			para1752{[]int{1, 1, 1}},
			ans1752{true},
		},

		{
			para1752{[]int{2, 1}},
			ans1752{true},
		},

		{
			para1752{[]int{1, 3, 2, 4}},
			ans1752{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1752------------------------\n")

	for _, q := range qs {
		_, p := q.ans1752, q.para1752
		fmt.Printf("【input】:%v       【output】:%v\n", p, check(p.nums))
	}
	fmt.Printf("\n\n\n")
}
