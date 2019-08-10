package leetcode

import (
	"fmt"
	"testing"
)

type question229 struct {
	para229
	ans229
}

// para 是参数
// one 代表第一个参数
type para229 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans229 struct {
	one []int
}

func Test_Problem229(t *testing.T) {

	qs := []question229{

		question229{
			para229{[]int{3, 2, 3}},
			ans229{[]int{3}},
		},

		question229{
			para229{[]int{1, 1, 1, 3, 3, 2, 2, 2}},
			ans229{[]int{1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 229------------------------\n")

	for _, q := range qs {
		_, p := q.ans229, q.para229
		fmt.Printf("【input】:%v       【output】:%v\n", p, majorityElement229(p.s))
	}
	fmt.Printf("\n\n\n")
}
