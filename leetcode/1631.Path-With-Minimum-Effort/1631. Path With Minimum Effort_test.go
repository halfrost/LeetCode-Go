package leetcode

import (
	"fmt"
	"testing"
)

type question1631 struct {
	para1631
	ans1631
}

// para 是参数
// one 代表第一个参数
type para1631 struct {
	heights [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1631 struct {
	one int
}

func Test_Problem1631(t *testing.T) {

	qs := []question1631{

		{
			para1631{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}},
			ans1631{2},
		},

		{
			para1631{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}},
			ans1631{1},
		},

		{
			para1631{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}},
			ans1631{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1631------------------------\n")

	for _, q := range qs {
		_, p := q.ans1631, q.para1631
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minimumEffortPath(p.heights))
	}
	fmt.Printf("\n\n\n")
}
