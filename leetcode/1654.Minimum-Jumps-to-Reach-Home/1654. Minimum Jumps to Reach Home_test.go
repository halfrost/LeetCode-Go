package leetcode

import (
	"fmt"
	"testing"
)

type question1654 struct {
	para1654
	ans1654
}

// para 是参数
// one 代表第一个参数
type para1654 struct {
	forbidden []int
	a         int
	b         int
	x         int
}

// ans 是答案
// one 代表第一个答案
type ans1654 struct {
	one int
}

func Test_Problem1654(t *testing.T) {

	qs := []question1654{

		{
			para1654{[]int{14, 4, 18, 1, 15}, 3, 15, 9},
			ans1654{3},
		},

		{
			para1654{[]int{8, 3, 16, 6, 12, 20}, 15, 13, 11},
			ans1654{-1},
		},

		{
			para1654{[]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7},
			ans1654{2},
		},

		{
			para1654{[]int{1998}, 1999, 2000, 2000},
			ans1654{3998},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1654------------------------\n")

	for _, q := range qs {
		_, p := q.ans1654, q.para1654
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minimumJumps(p.forbidden, p.a, p.b, p.x))
	}
	fmt.Printf("\n\n\n")
}
