package leetcode

import (
	"fmt"
	"testing"
)

type question1275 struct {
	para1275
	ans1275
}

// para 是参数
// one 代表第一个参数
type para1275 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1275 struct {
	one string
}

func Test_Problem1275(t *testing.T) {

	qs := []question1275{

		{
			para1275{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}},
			ans1275{"A"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}},
			ans1275{"B"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}},
			ans1275{"Draw"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}}},
			ans1275{"Pending"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1275------------------------\n")

	for _, q := range qs {
		_, p := q.ans1275, q.para1275
		fmt.Printf("【input】:%v       【output】:%v\n", p, tictactoe(p.one))
	}
	fmt.Printf("\n\n\n")
}
