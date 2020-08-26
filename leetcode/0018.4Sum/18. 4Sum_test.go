package leetcode

import (
	"fmt"
	"testing"
)

type question18 struct {
	para18
	ans18
}

// para 是参数
// one 代表第一个参数
type para18 struct {
	a []int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans18 struct {
	one [][]int
}

func Test_Problem18(t *testing.T) {

	qs := []question18{

		{
			para18{[]int{1, 1, 1, 1}, 4},
			ans18{[][]int{{1, 1, 1, 1}}},
		},

		{
			para18{[]int{0, 1, 5, 0, 1, 5, 5, -4}, 11},
			ans18{[][]int{{-4, 5, 5, 5}, {0, 1, 5, 5}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans18{[][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 0},
			ans18{[][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {0, 0, 0, 0}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 1},
			ans18{[][]int{{-1, 0, 0, 2}, {-2, 0, 1, 2}, {0, 0, 0, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 18------------------------\n")

	for _, q := range qs {
		_, p := q.ans18, q.para18
		fmt.Printf("【input】:%v       【output】:%v\n", p, fourSum(p.a, p.t))
	}
	fmt.Printf("\n\n\n")
}
