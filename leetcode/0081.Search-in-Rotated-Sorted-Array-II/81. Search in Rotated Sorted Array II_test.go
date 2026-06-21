package leetcode

import (
	"fmt"
	"testing"
)

type question81 struct {
	para81
	ans81
}

// para 是参数
// one 代表第一个参数
type para81 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans81 struct {
	one bool
}

func Test_Problem81(t *testing.T) {

	qs := []question81{

		{
			para81{[]int{2, 5, 6, 0, 0, 1, 2}, 0},
			ans81{true},
		},

		{
			para81{[]int{2, 5, 6, 0, 0, 1, 2}, 3},
			ans81{false},
		},

		{
			para81{[]int{}, 1},
			ans81{false},
		},

		{
			para81{[]int{1, 0, 1, 1, 1}, 0},
			ans81{true},
		},

		{
			para81{[]int{1, 1, 1, 0, 1}, 0},
			ans81{true},
		},

		{
			para81{[]int{4, 5, 6, 7, 0, 1, 2}, 6},
			ans81{true},
		},

		{
			para81{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans81{true},
		},

		{
			para81{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans81{false},
		},

		{
			para81{[]int{6, 7, 0, 1, 2}, 2},
			ans81{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 81------------------------\n")

	for _, q := range qs {
		a, p := q.ans81, q.para81
		got := search(p.nums, p.target)
		if got != a.one {
			t.Fatalf("input: %v target: %v, expected: %v, got: %v", p.nums, p.target, a.one, got)
		}
		fmt.Printf("【input】:%v    【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
