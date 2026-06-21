package leetcode

import (
	"fmt"
	"testing"
)

type question45 struct {
	para45
	ans45
}

// para 是参数
// one 代表第一个参数
type para45 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans45 struct {
	one int
}

func Test_Problem45(t *testing.T) {

	qs := []question45{

		{
			para45{[]int{2, 3, 1, 1, 4}},
			ans45{2},
		},

		{
			para45{[]int{2, 3, 0, 1, 4}},
			ans45{2},
		},

		{
			para45{[]int{0}},
			ans45{0},
		},

		{
			para45{[]int{}},
			ans45{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 45------------------------\n")

	for _, q := range qs {
		a, p := q.ans45, q.para45
		got := jump(p.nums)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
