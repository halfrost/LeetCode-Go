package leetcode

import (
	"fmt"
	"testing"
)

type question1203 struct {
	para1203
	ans1203
}

// para 是参数
// one 代表第一个参数
type para1203 struct {
	n           int
	m           int
	group       []int
	beforeItems [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1203 struct {
	one []int
}

func Test_Problem1203(t *testing.T) {

	qs := []question1203{

		{
			para1203{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}},
			ans1203{[]int{6, 3, 4, 5, 2, 0, 7, 1}},
		},

		{
			para1203{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}},
			ans1203{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1203------------------------\n")

	for _, q := range qs {
		_, p := q.ans1203, q.para1203
		fmt.Printf("【input】:%v       【output】:%v\n", p, sortItems1(p.n, p.m, p.group, p.beforeItems))
	}
	fmt.Printf("\n\n\n")
}
