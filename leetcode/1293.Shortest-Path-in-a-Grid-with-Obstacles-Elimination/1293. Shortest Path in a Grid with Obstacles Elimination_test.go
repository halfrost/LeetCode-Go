package leetcode

import (
	"fmt"
	"testing"
)

type question1293 struct {
	para1293
	ans1293
}

// para 是参数
// one 代表第一个参数
type para1293 struct {
	grid [][]int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1293 struct {
	one int
}

func Test_Problem1293(t *testing.T) {

	qs := []question1293{

		{
			para1293{[][]int{
				{0, 0, 0},
			}, 1},
			ans1293{2},
		},

		{
			para1293{[][]int{
				{0, 1, 1}, {0, 1, 1}, {0, 0, 0}, {0, 1, 0}, {0, 1, 0},
			}, 2},
			ans1293{6},
		},

		{
			para1293{[][]int{
				{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0},
			}, 1},
			ans1293{6},
		},

		{
			para1293{[][]int{
				{0, 1, 1}, {1, 1, 1}, {1, 0, 0},
			}, 1},
			ans1293{-1},
		},

		{
			para1293{[][]int{
				{0},
			}, 1},
			ans1293{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1293------------------------\n")

	for _, q := range qs {
		_, p := q.ans1293, q.para1293
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestPath(p.grid, p.k))
	}
	fmt.Printf("\n\n\n")
}
