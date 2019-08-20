package leetcode

import (
	"fmt"
	"testing"
)

type question491 struct {
	para491
	ans491
}

// para 是参数
// one 代表第一个参数
type para491 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans491 struct {
	one [][]int
}

func Test_Problem491(t *testing.T) {

	qs := []question491{

		question491{
			para491{[]int{4, 3, 2, 1}},
			ans491{[][]int{}},
		},

		question491{
			para491{[]int{4, 6, 7, 7}},
			ans491{[][]int{[]int{4, 6}, []int{4, 7}, []int{4, 6, 7}, []int{4, 6, 7, 7}, []int{6, 7}, []int{6, 7, 7}, []int{7, 7}, []int{4, 7, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 491------------------------\n")

	for _, q := range qs {
		_, p := q.ans491, q.para491
		fmt.Printf("【input】:%v       【output】:%v\n", p, findSubsequences(p.one))
	}
	fmt.Printf("\n\n\n")
}
