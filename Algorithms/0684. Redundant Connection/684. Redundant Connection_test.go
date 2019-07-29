package leetcode

import (
	"fmt"
	"testing"
)

type question684 struct {
	para684
	ans684
}

// para 是参数
// one 代表第一个参数
type para684 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans684 struct {
	one []int
}

func Test_Problem684(t *testing.T) {

	qs := []question684{

		question684{
			para684{[][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}},
			ans684{[]int{2, 3}},
		},

		question684{
			para684{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 4}, []int{1, 5}}},
			ans684{[]int{1, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 684------------------------\n")

	for _, q := range qs {
		_, p := q.ans684, q.para684
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRedundantConnection(p.one))
	}
	fmt.Printf("\n\n\n")
}
