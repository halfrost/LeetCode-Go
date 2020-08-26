package leetcode

import (
	"fmt"
	"testing"
)

type question850 struct {
	para850
	ans850
}

// para 是参数
// one 代表第一个参数
type para850 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans850 struct {
	one int
}

func Test_Problem850(t *testing.T) {

	qs := []question850{

		{
			para850{[][]int{{0, 0, 3, 3}, {2, 0, 5, 3}, {1, 1, 4, 4}}},
			ans850{18},
		},

		{
			para850{[][]int{{0, 0, 1, 1}, {2, 2, 3, 3}}},
			ans850{2},
		},

		{
			para850{[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}},
			ans850{6},
		},

		{
			para850{[][]int{{0, 0, 1000000000, 1000000000}}},
			ans850{49},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 850------------------------\n")

	for _, q := range qs {
		_, p := q.ans850, q.para850
		fmt.Printf("【input】:%v       【output】:%v\n", p, rectangleArea(p.one))
	}
	fmt.Printf("\n\n\n")
}
