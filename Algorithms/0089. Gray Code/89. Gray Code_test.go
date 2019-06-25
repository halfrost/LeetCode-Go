package leetcode

import (
	"fmt"
	"testing"
)

type question89 struct {
	para89
	ans89
}

// para 是参数
// one 代表第一个参数
type para89 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans89 struct {
	one []int
}

func Test_Problem89(t *testing.T) {

	qs := []question89{

		question89{
			para89{2},
			ans89{[]int{0, 1, 3, 2}},
		},

		question89{
			para89{0},
			ans89{[]int{0}},
		},

		question89{
			para89{3},
			ans89{[]int{0, 1, 3, 2, 6, 7, 5, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 89------------------------\n")

	for _, q := range qs {
		_, p := q.ans89, q.para89
		fmt.Printf("【input】:%v       【output】:%v\n", p, grayCode(p.one))
	}
	fmt.Printf("\n\n\n")
}
