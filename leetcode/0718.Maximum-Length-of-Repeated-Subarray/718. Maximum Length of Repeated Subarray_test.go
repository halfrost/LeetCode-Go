package leetcode

import (
	"fmt"
	"testing"
)

type question718 struct {
	para718
	ans718
}

// para 是参数
// one 代表第一个参数
type para718 struct {
	A []int
	B []int
}

// ans 是答案
// one 代表第一个答案
type ans718 struct {
	one int
}

func Test_Problem718(t *testing.T) {

	qs := []question718{

		question718{
			para718{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
			ans718{5},
		},

		question718{
			para718{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}},
			ans718{3},
		},

		question718{
			para718{[]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}},
			ans718{4},
		},

		question718{
			para718{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}},
			ans718{59},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 718------------------------\n")

	for _, q := range qs {
		_, p := q.ans718, q.para718
		fmt.Printf("【input】:%v       【output】:%v\n", p, findLength(p.A, p.B))
	}
	fmt.Printf("\n\n\n")
}
