package leetcode

import (
	"fmt"
	"testing"
)

type question977 struct {
	para977
	ans977
}

// para 是参数
// one 代表第一个参数
type para977 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans977 struct {
	one []int
}

func Test_Problem977(t *testing.T) {

	qs := []question977{

		question977{
			para977{[]int{-4, -1, 0, 3, 10}},
			ans977{[]int{0, 1, 9, 16, 100}},
		},

		question977{
			para977{[]int{1}},
			ans977{[]int{1}},
		},

		question977{
			para977{[]int{-7, -3, 2, 3, 11}},
			ans977{[]int{4, 9, 9, 49, 121}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 977------------------------\n")

	for _, q := range qs {
		_, p := q.ans977, q.para977
		fmt.Printf("【input】:%v       【output】:%v\n", p, sortedSquares(p.one))
	}
	fmt.Printf("\n\n\n")
}
