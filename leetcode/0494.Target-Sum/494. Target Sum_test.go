package leetcode

import (
	"fmt"
	"testing"
)

type question494 struct {
	para494
	ans494
}

// para 是参数
// one 代表第一个参数
type para494 struct {
	nums []int
	S    int
}

// ans 是答案
// one 代表第一个答案
type ans494 struct {
	one int
}

func Test_Problem494(t *testing.T) {

	qs := []question494{

		{
			para494{[]int{1, 1, 1, 1, 1}, 3},
			ans494{5},
		},
		{
			para494{[]int{1}, 2},
			ans494{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 494------------------------\n")

	for _, q := range qs {
		a, p := q.ans494, q.para494
		got := findTargetSumWays(p.nums, p.S)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("findTargetSumWays(%v, %v) = %v, want %v", p.nums, p.S, got, a.one)
		}
		if got1 := findTargetSumWays1(p.nums, p.S); got1 != a.one {
			t.Fatalf("findTargetSumWays1(%v, %v) = %v, want %v", p.nums, p.S, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
