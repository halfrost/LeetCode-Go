package leetcode

import (
	"fmt"
	"testing"
)

type question1725 struct {
	para1725
	ans1725
}

// para 是参数
// one 代表第一个参数
type para1725 struct {
	rectangles [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1725 struct {
	one int
}

func Test_Problem1725(t *testing.T) {

	qs := []question1725{

		{
			para1725{[][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}},
			ans1725{3},
		},

		{
			para1725{[][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}},
			ans1725{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1725------------------------\n")

	for _, q := range qs {
		_, p := q.ans1725, q.para1725
		fmt.Printf("【input】:%v       【output】:%v\n", p, countGoodRectangles(p.rectangles))
	}
	fmt.Printf("\n\n\n")
}
