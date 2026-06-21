package leetcode

import (
	"fmt"
	"testing"
)

type question674 struct {
	para674
	ans674
}

// para 是参数
// one 代表第一个参数
type para674 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans674 struct {
	one int
}

func Test_Problem674(t *testing.T) {

	qs := []question674{

		{
			para674{[]int{1, 3, 5, 4, 7}},
			ans674{3},
		},

		{
			para674{[]int{2, 2, 2, 2, 2}},
			ans674{1},
		},

		{
			para674{[]int{1, 3, 5, 7}},
			ans674{4},
		},

		{
			para674{[]int{}},
			ans674{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 674------------------------\n")

	for _, q := range qs {
		a, p := q.ans674, q.para674
		got := findLengthOfLCIS(p.nums)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input %v: got %v, want %v", p.nums, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
