package leetcode

import (
	"fmt"
	"testing"
)

type question812 struct {
	para812
	ans812
}

// para 是参数
// one 代表第一个参数
type para812 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans812 struct {
	one float64
}

func Test_Problem812(t *testing.T) {

	qs := []question812{

		{
			para812{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			ans812{2.0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 812------------------------\n")

	for _, q := range qs {
		_, p := q.ans812, q.para812
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestTriangleArea(p.one))
	}
	fmt.Printf("\n\n\n")
}
