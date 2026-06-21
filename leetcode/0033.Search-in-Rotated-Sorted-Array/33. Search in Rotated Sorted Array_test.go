package leetcode

import (
	"fmt"
	"testing"
)

type question33 struct {
	para33
	ans33
}

// para 是参数
// one 代表第一个参数
type para33 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans33 struct {
	one int
}

func Test_Problem33(t *testing.T) {

	qs := []question33{

		{
			para33{[]int{3, 1}, 1},
			ans33{1},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans33{4},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans33{-1},
		},

		{
			para33{[]int{5, 6, 7, 0, 1, 2, 3, 4}, 2},
			ans33{5},
		},

		{
			para33{[]int{5, 6, 7, 0, 1, 2, 3, 4}, 6},
			ans33{1},
		},

		{
			para33{[]int{1, 1, 1, 1, 1, 1, 1}, 2},
			ans33{-1},
		},

		{
			para33{[]int{}, 5},
			ans33{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 33------------------------\n")

	for _, q := range qs {
		a, p := q.ans33, q.para33
		got := search33(p.nums, p.target)
		fmt.Printf("【input】:%v    【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
