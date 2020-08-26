package leetcode

import (
	"fmt"
	"testing"
)

type question1266 struct {
	para1266
	ans1266
}

// para 是参数
// one 代表第一个参数
type para1266 struct {
	points [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1266 struct {
	one int
}

func Test_Problem1266(t *testing.T) {

	qs := []question1266{

		{
			para1266{[][]int{{1, 1}, {3, 4}, {-1, 0}}},
			ans1266{7},
		},

		{
			para1266{[][]int{{3, 2}, {-2, 2}}},
			ans1266{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1266------------------------\n")

	for _, q := range qs {
		_, p := q.ans1266, q.para1266
		fmt.Printf("【input】:%v       【output】:%v\n", p, minTimeToVisitAllPoints(p.points))
	}
	fmt.Printf("\n\n\n")
}
