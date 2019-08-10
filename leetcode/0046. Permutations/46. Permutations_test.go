package leetcode

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	one [][]int
}

func Test_Problem46(t *testing.T) {

	qs := []question46{

		question46{
			para46{[]int{1, 2, 3}},
			ans46{[][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, q := range qs {
		_, p := q.ans46, q.para46
		fmt.Printf("【input】:%v       【output】:%v\n", p, permute(p.s))
	}
	fmt.Printf("\n\n\n")
}
