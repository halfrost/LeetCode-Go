package leetcode

import (
	"fmt"
	"testing"
)

type question946 struct {
	para946
	ans946
}

// para 是参数
// one 代表第一个参数
type para946 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans946 struct {
	one bool
}

func Test_Problem946(t *testing.T) {

	qs := []question946{
		question946{
			para946{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
			ans946{true},
		},

		question946{
			para946{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}},
			ans946{false},
		},

		question946{
			para946{[]int{1, 0}, []int{1, 0}},
			ans946{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 946------------------------\n")

	for _, q := range qs {
		_, p := q.ans946, q.para946
		fmt.Printf("【input】:%v       【output】:%v\n", p, validateStackSequences(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
