package leetcode

import (
	"fmt"
	"testing"
)

type question2021 struct {
	para2021
	ans2021
}

// para 是参数
type para2021 struct {
	lights [][]int
}

// ans 是答案
type ans2021 struct {
	ans int
}

func Test_Problem2021(t *testing.T) {

	qs := []question2021{

		{
			para2021{[][]int{{-3, 2}, {1, 2}, {3, 3}}},
			ans2021{-1},
		},

		{
			para2021{[][]int{{1, 0}, {0, 1}}},
			ans2021{1},
		},

		{
			para2021{[][]int{{1, 2}}},
			ans2021{-1},
		},

		{
			para2021{[][]int{{1, 1}, {2, 4}, {-1, 0}, {-3, 5}, {1, 2}}},
			ans2021{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2021------------------------\n")

	for _, q := range qs {
		_, p := q.ans2021, q.para2021
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", brightestPosition(p.lights))
	}
	fmt.Printf("\n\n\n")
}
