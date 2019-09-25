package leetcode

import (
	"fmt"
	"testing"
)

type question74 struct {
	para74
	ans74
}

// para 是参数
// one 代表第一个参数
type para74 struct {
	matrix [][]int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans74 struct {
	one bool
}

func Test_Problem74(t *testing.T) {

	qs := []question74{

		question74{
			para74{[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}, 3},
			ans74{true},
		},

		question74{
			para74{[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}, 13},
			ans74{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 74------------------------\n")

	for _, q := range qs {
		_, p := q.ans74, q.para74
		fmt.Printf("【input】:%v       【output】:%v\n", p, searchMatrix(p.matrix, p.target))
	}
	fmt.Printf("\n\n\n")
}
