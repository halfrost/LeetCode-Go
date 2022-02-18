package leetcode

import (
	"fmt"
	"testing"
)

type question1791 struct {
	para1791
	ans1791
}

// para 是参数
type para1791 struct {
	edges [][]int
}

// ans 是答案
type ans1791 struct {
	ans int
}

func Test_Problem1791(t *testing.T) {

	qs := []question1791{

		{
			para1791{[][]int{{1, 2}, {2, 3}, {4, 2}}},
			ans1791{2},
		},

		{
			para1791{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}},
			ans1791{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1791------------------------\n")

	for _, q := range qs {
		_, p := q.ans1791, q.para1791
		fmt.Printf("【input】:%v      ", p.edges)
		fmt.Printf("【output】:%v      \n", findCenter(p.edges))
	}
	fmt.Printf("\n\n\n")
}
