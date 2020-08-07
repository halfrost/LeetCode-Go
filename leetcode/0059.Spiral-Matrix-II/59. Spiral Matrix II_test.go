package leetcode

import (
	"fmt"
	"testing"
)

type question59 struct {
	para59
	ans59
}

// para 是参数
// one 代表第一个参数
type para59 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans59 struct {
	one [][]int
}

func Test_Problem59(t *testing.T) {

	qs := []question59{

		question59{
			para59{3},
			ans59{[][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}},
		},

		question59{
			para59{4},
			ans59{[][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 59------------------------\n")

	for _, q := range qs {
		_, p := q.ans59, q.para59
		fmt.Printf("【input】:%v       【output】:%v\n", p, generateMatrix(p.one))
	}
	fmt.Printf("\n\n\n")
}
