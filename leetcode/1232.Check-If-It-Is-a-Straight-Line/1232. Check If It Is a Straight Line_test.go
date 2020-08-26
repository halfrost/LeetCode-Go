package leetcode

import (
	"fmt"
	"testing"
)

type question1232 struct {
	para1232
	ans1232
}

// para 是参数
// one 代表第一个参数
type para1232 struct {
	arr [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1232 struct {
	one bool
}

func Test_Problem1232(t *testing.T) {

	qs := []question1232{

		{
			para1232{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}},
			ans1232{true},
		},

		{
			para1232{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}},
			ans1232{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1232------------------------\n")

	for _, q := range qs {
		_, p := q.ans1232, q.para1232
		fmt.Printf("【input】:%v       【output】:%v\n", p, checkStraightLine(p.arr))
	}
	fmt.Printf("\n\n\n")
}
