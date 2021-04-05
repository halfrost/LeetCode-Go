package leetcode

import (
	"fmt"
	"testing"
)

type question775 struct {
	para775
	ans775
}

// para 是参数
// one 代表第一个参数
type para775 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans775 struct {
	one bool
}

func Test_Problem775(t *testing.T) {

	qs := []question775{
		{
			para775{[]int{1, 0, 2}},
			ans775{true},
		},

		{
			para775{[]int{1, 2, 0}},
			ans775{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 775------------------------\n")

	for _, q := range qs {
		_, p := q.ans775, q.para775
		fmt.Printf("【input】:%v       【output】:%v\n", p, isIdealPermutation(p.A))
	}
	fmt.Printf("\n\n\n")
}
