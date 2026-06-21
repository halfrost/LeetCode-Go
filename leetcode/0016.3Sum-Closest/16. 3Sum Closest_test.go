package leetcode

import (
	"fmt"
	"testing"
)

type question16 struct {
	para16
	ans16
}

// para 是参数
// one 代表第一个参数
type para16 struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans16 struct {
	one int
}

func Test_Problem16(t *testing.T) {

	qs := []question16{

		{
			para16{[]int{-1, 0, 1, 1, 55}, 3},
			ans16{2},
		},

		{
			para16{[]int{0, 0, 0}, 1},
			ans16{0},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},

		{
			para16{[]int{1, 1, -1}, 0},
			ans16{1},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},

		// exact match triggers early return (sum == target)
		{
			para16{[]int{0, 1, 2, 3}, 6},
			ans16{6},
		},

		// duplicate leading values trigger the i>0 && nums[i]==nums[i-1] continue branch
		{
			para16{[]int{1, 1, 1, 0, 5}, 100},
			ans16{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 16------------------------\n")

	for _, q := range qs {
		a, p := q.ans16, q.para16
		out := threeSumClosest(append([]int{}, p.a...), p.target)
		fmt.Printf("【input】:%v       【output】:%v\n", p, out)
		if out != a.one {
			t.Fatalf("threeSumClosest(%v, %d) = %d, want %d", p.a, p.target, out, a.one)
		}
		if out2 := threeSumClosest1(append([]int{}, p.a...), p.target); out2 != a.one {
			t.Fatalf("threeSumClosest1(%v, %d) = %d, want %d", p.a, p.target, out2, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
