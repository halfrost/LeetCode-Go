package leetcode

import (
	"fmt"
	"testing"
)

type question64 struct {
	para64
	ans64
}

// para 是参数
// one 代表第一个参数
type para64 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans64 struct {
	one int
}

func cloneGrid64(g [][]int) [][]int {
	c := make([][]int, len(g))
	for i := range g {
		c[i] = make([]int, len(g[i]))
		copy(c[i], g[i])
	}
	return c
}

func Test_Problem64(t *testing.T) {

	qs := []question64{

		{
			para64{[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			ans64{7},
		},

		{
			para64{[][]int{}},
			ans64{0},
		},

		{
			para64{[][]int{{}}},
			ans64{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 64------------------------\n")

	for _, q := range qs {
		a, p := q.ans64, q.para64
		// minPathSum1 covers the empty/zero-size guard branches
		got1 := minPathSum1(cloneGrid64(p.og))
		if got1 != a.one {
			t.Fatalf("minPathSum1(%v) = %d, want %d", p.og, got1, a.one)
		}
		// minPathSum panics on empty grid, so only call it on non-empty input
		if len(p.og) > 0 && len(p.og[0]) > 0 {
			got := minPathSum(cloneGrid64(p.og))
			if got != a.one {
				t.Fatalf("minPathSum(%v) = %d, want %d", p.og, got, a.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		} else {
			fmt.Printf("【input】:%v       【output】:%v\n", p, got1)
		}
	}
	fmt.Printf("\n\n\n")
}
