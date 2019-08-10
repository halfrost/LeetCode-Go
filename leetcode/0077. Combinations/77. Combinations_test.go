package leetcode

import (
	"fmt"
	"testing"
)

type question77 struct {
	para77
	ans77
}

// para 是参数
// one 代表第一个参数
type para77 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans77 struct {
	one [][]int
}

func Test_Problem77(t *testing.T) {

	qs := []question77{

		question77{
			para77{4, 2},
			ans77{[][]int{[]int{2, 4}, []int{3, 4}, []int{2, 3}, []int{1, 2}, []int{1, 3}, []int{1, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 77------------------------\n")

	for _, q := range qs {
		_, p := q.ans77, q.para77
		fmt.Printf("【input】:%v       【output】:%v\n", p, combine(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
