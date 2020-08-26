package leetcode

import (
	"fmt"
	"testing"
)

type question1292 struct {
	para1292
	ans1292
}

// para 是参数
// one 代表第一个参数
type para1292 struct {
	mat       [][]int
	threshold int
}

// ans 是答案
// one 代表第一个答案
type ans1292 struct {
	one int
}

func Test_Problem1292(t *testing.T) {

	qs := []question1292{

		{
			para1292{[][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4},
			ans1292{2},
		},

		{
			para1292{[][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 1},
			ans1292{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1292------------------------\n")

	for _, q := range qs {
		_, p := q.ans1292, q.para1292
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxSideLength(p.mat, p.threshold))
	}
	fmt.Printf("\n\n\n")
}
