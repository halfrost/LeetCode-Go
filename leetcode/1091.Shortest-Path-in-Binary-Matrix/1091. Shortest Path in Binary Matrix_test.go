package leetcode

import (
	"fmt"
	"testing"
)

type question1091 struct {
	para1091
	ans1091
}

// para 是参数
// one 代表第一个参数
type para1091 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1091 struct {
	one int
}

func Test_Problem1091(t *testing.T) {

	qs := []question1091{

		{
			para1091{[][]int{{0, 1}, {1, 0}}},
			ans1091{2},
		},

		{
			para1091{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}},
			ans1091{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1091------------------------\n")

	for _, q := range qs {
		_, p := q.ans1091, q.para1091
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestPathBinaryMatrix(p.grid))

	}
	fmt.Printf("\n\n\n")
}
