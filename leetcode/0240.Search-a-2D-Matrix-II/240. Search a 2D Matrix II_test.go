package leetcode

import (
	"fmt"
	"testing"
)

type question240 struct {
	para240
	ans240
}

// para 是参数
// one 代表第一个参数
type para240 struct {
	matrix [][]int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans240 struct {
	one bool
}

func Test_Problem240(t *testing.T) {

	qs := []question240{

		{
			para240{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5},
			ans240{true},
		},

		{
			para240{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20},
			ans240{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 240------------------------\n")

	for _, q := range qs {
		_, p := q.ans240, q.para240
		fmt.Printf("【input】:%v       【output】:%v\n", p, searchMatrix240(p.matrix, p.target))
	}
	fmt.Printf("\n\n\n")
}
