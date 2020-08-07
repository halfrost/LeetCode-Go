package leetcode

import (
	"fmt"
	"testing"
)

type question84 struct {
	para84
	ans84
}

// para 是参数
// one 代表第一个参数
type para84 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans84 struct {
	one int
}

func Test_Problem84(t *testing.T) {

	qs := []question84{

		question84{
			para84{[]int{2, 1, 5, 6, 2, 3}},
			ans84{10},
		},

		question84{
			para84{[]int{1}},
			ans84{1},
		},

		question84{
			para84{[]int{1, 1}},
			ans84{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 84------------------------\n")

	for _, q := range qs {
		_, p := q.ans84, q.para84
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestRectangleArea(p.one))
	}
	fmt.Printf("\n\n\n")
}
