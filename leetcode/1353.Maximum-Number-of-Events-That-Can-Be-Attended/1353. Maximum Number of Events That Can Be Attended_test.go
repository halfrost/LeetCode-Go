package leetcode

import (
	"fmt"
	"testing"
)

type question1353 struct {
	para1353
	ans1353
}

// para 是参数
// one 代表第一个参数
type para1353 struct {
	events [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1353 struct {
	one int
}

func Test_Problem1353(t *testing.T) {

	qs := []question1353{

		{
			para1353{[][]int{{1, 2}, {2, 3}, {3, 4}}},
			ans1353{3},
		},

		{
			para1353{[][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}},
			ans1353{4},
		},

		{
			para1353{[][]int{{1, 100000}}},
			ans1353{1},
		},

		{
			para1353{[][]int{{1, 1}, {2, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}},
			ans1353{7},
		},

		{
			para1353{[][]int{{1, 2}, {2, 2}, {3, 3}, {3, 4}, {3, 4}}},
			ans1353{4},
		},

		{
			para1353{[][]int{{1, 10}, {2, 2}, {2, 2}, {2, 2}, {2, 2}}},
			ans1353{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1353------------------------\n")

	for _, q := range qs {
		_, p := q.ans1353, q.para1353
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxEvents(p.events))
	}
	fmt.Printf("\n\n\n")
}
