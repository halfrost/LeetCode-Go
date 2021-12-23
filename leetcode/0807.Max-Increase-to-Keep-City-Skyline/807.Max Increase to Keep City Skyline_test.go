package leetcode

import (
	"fmt"
	"testing"
)

type question807 struct {
	para807
	ans807
}

// para 是参数
type para807 struct {
	grid [][]int
}

// ans 是答案
type ans807 struct {
	ans int
}

func Test_Problem807(t *testing.T) {

	qs := []question807{

		{
			para807{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}},
			ans807{35},
		},

		{
			para807{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			ans807{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 807------------------------\n")

	for _, q := range qs {
		_, p := q.ans807, q.para807
		fmt.Printf("【input】:%v    【output】:%v\n", p.grid, maxIncreaseKeepingSkyline(p.grid))
	}
	fmt.Printf("\n\n\n")
}
