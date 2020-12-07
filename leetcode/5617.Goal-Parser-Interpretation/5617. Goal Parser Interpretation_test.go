package leetcode

import (
	"fmt"
	"testing"
)

type question491 struct {
	para491
	ans491
}

// para 是参数
// one 代表第一个参数
type para491 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans491 struct {
	one []int
}

func Test_Problem491(t *testing.T) {

	qs := []question491{

		{
			para491{[]int{3, 5, 2, 6}, 2},
			ans491{[]int{2, 6}},
		},

		{
			para491{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4},
			ans491{[]int{2, 3, 3, 4}},
		},

		{
			para491{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4},
			ans491{[]int{2, 3, 3, 4}},
		},

		{
			para491{[]int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}, 3},
			ans491{[]int{8, 80, 2}},
		},

		{
			para491{[]int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, 24},
			ans491{[]int{10, 23, 61, 62, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 491------------------------\n")

	for _, q := range qs {
		_, p := q.ans491, q.para491
		fmt.Printf("【input】:%v       【output】:%v\n", p, mostCompetitive(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
