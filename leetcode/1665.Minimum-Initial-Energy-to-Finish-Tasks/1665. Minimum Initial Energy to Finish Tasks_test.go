package leetcode

import (
	"fmt"
	"testing"
)

type question1665 struct {
	para1665
	ans1665
}

// para 是参数
// one 代表第一个参数
type para1665 struct {
	tasks [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1665 struct {
	one int
}

func Test_Problem1665(t *testing.T) {

	qs := []question1665{

		{
			para1665{[][]int{{1, 2}, {2, 4}, {4, 8}}},
			ans1665{8},
		},

		{
			para1665{[][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}},
			ans1665{32},
		},

		{
			para1665{[][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}},
			ans1665{27},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1665------------------------\n")

	for _, q := range qs {
		_, p := q.ans1665, q.para1665
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minimumEffort(p.tasks))
	}
	fmt.Printf("\n\n\n")
}
