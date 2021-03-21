package leetcode

import (
	"fmt"
	"testing"
)

type question1673 struct {
	para1673
	ans1673
}

// para 是参数
// one 代表第一个参数
type para1673 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1673 struct {
	one []int
}

func Test_Problem1673(t *testing.T) {

	qs := []question1673{

		{
			para1673{[]int{3, 5, 2, 6}, 2},
			ans1673{[]int{2, 6}},
		},

		{
			para1673{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4},
			ans1673{[]int{2, 3, 3, 4}},
		},

		{
			para1673{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4},
			ans1673{[]int{2, 3, 3, 4}},
		},

		{
			para1673{[]int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}, 3},
			ans1673{[]int{8, 80, 2}},
		},

		{
			para1673{[]int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, 24},
			ans1673{[]int{10, 23, 61, 62, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1673------------------------\n")

	for _, q := range qs {
		_, p := q.ans1673, q.para1673
		fmt.Printf("【input】:%v       【output】:%v\n", p, mostCompetitive(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
