package leetcode

import (
	"fmt"
	"testing"
)

type question33 struct {
	para33
	ans33
}

// para 是参数
// one 代表第一个参数
type para33 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans33 struct {
	one int
}

func Test_Problem33(t *testing.T) {

	qs := []question33{

		question33{
			para33{[]int{3, 1}, 1},
			ans33{1},
		},

		question33{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans33{4},
		},

		question33{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans33{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 33------------------------\n")

	for _, q := range qs {
		_, p := q.ans33, q.para33
		fmt.Printf("【input】:%v    【output】:%v\n", p, search33(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
