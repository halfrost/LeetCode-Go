package leetcode

import (
	"fmt"
	"testing"
)

type question766 struct {
	para766
	ans766
}

// para 是参数
// one 代表第一个参数
type para766 struct {
	A [][]int
}

// ans 是答案
// one 代表第一个答案
type ans766 struct {
	B bool
}

func Test_Problem766(t *testing.T) {

	qs := []question766{

		question766{
			para766{[][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}}},
			ans766{true},
		},

		question766{
			para766{[][]int{[]int{1, 2}, []int{2, 2}}},
			ans766{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 766------------------------\n")

	for _, q := range qs {
		_, p := q.ans766, q.para766
		fmt.Printf("【input】:%v       【output】:%v\n", p, isToeplitzMatrix(p.A))
	}
	fmt.Printf("\n\n\n")
}
