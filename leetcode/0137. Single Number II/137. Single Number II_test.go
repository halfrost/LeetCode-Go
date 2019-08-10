package leetcode

import (
	"fmt"
	"testing"
)

type question137 struct {
	para137
	ans137
}

// para 是参数
// one 代表第一个参数
type para137 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans137 struct {
	one int
}

func Test_Problem137(t *testing.T) {

	qs := []question137{

		question137{
			para137{[]int{2, 2, 3, 2}},
			ans137{3},
		},

		question137{
			para137{[]int{0, 1, 0, 1, 0, 1, 99}},
			ans137{99},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 137------------------------\n")

	for _, q := range qs {
		_, p := q.ans137, q.para137
		fmt.Printf("【input】:%v       【output】:%v\n", p, singleNumberII(p.s))
	}
	fmt.Printf("\n\n\n")
}
