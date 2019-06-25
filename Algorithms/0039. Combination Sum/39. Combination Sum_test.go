package leetcode

import (
	"fmt"
	"testing"
)

type question39 struct {
	para39
	ans39
}

// para 是参数
// one 代表第一个参数
type para39 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans39 struct {
	one [][]int
}

func Test_Problem39(t *testing.T) {

	qs := []question39{

		question39{
			para39{[]int{2, 3, 6, 7}, 7},
			ans39{[][]int{[]int{7}, []int{2, 2, 3}}},
		},

		question39{
			para39{[]int{2, 3, 5}, 8},
			ans39{[][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 39------------------------\n")

	for _, q := range qs {
		_, p := q.ans39, q.para39
		fmt.Printf("【input】:%v       【output】:%v\n", p, combinationSum(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
