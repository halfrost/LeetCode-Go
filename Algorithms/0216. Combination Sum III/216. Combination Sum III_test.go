package leetcode

import (
	"fmt"
	"testing"
)

type question216 struct {
	para216
	ans216
}

// para 是参数
// one 代表第一个参数
type para216 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans216 struct {
	one [][]int
}

func Test_Problem216(t *testing.T) {

	qs := []question216{

		question216{
			para216{3, 7},
			ans216{[][]int{[]int{1, 2, 4}}},
		},
		question216{
			para216{3, 9},
			ans216{[][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 216------------------------\n")

	for _, q := range qs {
		_, p := q.ans216, q.para216
		fmt.Printf("【input】:%v       【output】:%v\n", p, combinationSum3(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
