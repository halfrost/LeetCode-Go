package leetcode

import (
	"fmt"
	"testing"
)

type question918 struct {
	para918
	ans918
}

// para 是参数
// one 代表第一个参数
type para918 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans918 struct {
	one int
}

func Test_Problem918(t *testing.T) {

	qs := []question918{

		question918{
			para918{[]int{1, -2, 3, -2}},
			ans918{3},
		},

		question918{
			para918{[]int{5, -3, 5}},
			ans918{10},
		},

		question918{
			para918{[]int{3, -1, 2, -1}},
			ans918{4},
		},

		question918{
			para918{[]int{3, -2, 2, -3}},
			ans918{3},
		},

		question918{
			para918{[]int{-2, -3, -1}},
			ans918{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 918------------------------\n")

	for _, q := range qs {
		_, p := q.ans918, q.para918
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxSubarraySumCircular(p.one))
	}
	fmt.Printf("\n\n\n")
}
