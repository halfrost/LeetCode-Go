package leetcode

import (
	"fmt"
	"testing"
)

type question1260 struct {
	para1260
	ans1260
}

// para 是参数
// one 代表第一个参数
type para1260 struct {
	grid [][]int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1260 struct {
	one [][]int
}

func Test_Problem1260(t *testing.T) {

	qs := []question1260{

		{
			para1260{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, 2},
			ans1260{[][]int{{16, 17, 3}, {7, 8, 9}, {11, 13, 15}}},
		},

		{
			para1260{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, 10},
			ans1260{[][]int{{4, 2, 9, 3}, {8, 7, 15, 16}, {17, 12, 1, 10}}},
		},

		{
			para1260{[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4},
			ans1260{[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}}},
		},

		{
			para1260{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9},
			ans1260{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1260------------------------\n")

	for _, q := range qs {
		_, p := q.ans1260, q.para1260
		fmt.Printf("【input】:%v       【output】:%v\n", p, shiftGrid(p.grid, p.k))
	}
	fmt.Printf("\n\n\n")
}
