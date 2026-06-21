package leetcode

import (
	"fmt"
	"testing"
)

type question63 struct {
	para63
	ans63
}

// para 是参数
// one 代表第一个参数
type para63 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans63 struct {
	one int
}

func Test_Problem63(t *testing.T) {

	qs := []question63{

		{
			para63{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans63{2},
		},

		{
			para63{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{}},
			ans63{0},
		},

		{
			para63{[][]int{
				{1, 0},
				{0, 0},
			}},
			ans63{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 63------------------------\n")

	for _, q := range qs {
		a, p := q.ans63, q.para63
		got := uniquePathsWithObstacles(p.og)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
