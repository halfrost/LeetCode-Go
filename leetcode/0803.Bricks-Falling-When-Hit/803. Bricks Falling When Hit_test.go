package leetcode

import (
	"fmt"
	"testing"
)

type question803 struct {
	para803
	ans803
}

// para 是参数
// one 代表第一个参数
type para803 struct {
	grid [][]int
	hits [][]int
}

// ans 是答案
// one 代表第一个答案
type ans803 struct {
	one []int
}

func Test_Problem803(t *testing.T) {

	qs := []question803{

		{
			para803{[][]int{{1, 1, 1, 0, 1, 1, 1, 1}, {1, 0, 0, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 1, 1}, {1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}},
				[][]int{{4, 6}, {3, 0}, {2, 3}, {2, 6}, {4, 1}, {5, 2}, {2, 1}}},
			ans803{[]int{0, 2, 0, 0, 0, 0, 2}},
		},

		{
			para803{[][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}}},
			ans803{[]int{0, 3, 0}},
		},
		{
			para803{[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}},
			ans803{[]int{2}},
		},

		{
			para803{[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}, [][]int{{1, 1}, {1, 0}}},
			ans803{[]int{0, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 803------------------------\n")

	for _, q := range qs {
		_, p := q.ans803, q.para803
		fmt.Printf("【input】:%v       【output】:%v\n", p, hitBricks(p.grid, p.hits))
	}
	fmt.Printf("\n\n\n")
}
