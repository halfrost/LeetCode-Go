package leetcode

import (
	"fmt"
	"testing"
)

type question413 struct {
	para413
	ans413
}

// para 是参数
// one 代表第一个参数
type para413 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans413 struct {
	one int
}

func Test_Problem413(t *testing.T) {

	qs := []question413{

		{
			para413{[]int{1, 2, 3, 4}},
			ans413{3},
		},

		{
			para413{[]int{1, 2, 3, 4, 9}},
			ans413{3},
		},

		{
			para413{[]int{1, 2, 3, 4, 5, 6, 7}},
			ans413{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 413------------------------\n")

	for _, q := range qs {
		_, p := q.ans413, q.para413
		fmt.Printf("【input】:%v       【output】:%v\n", p, numberOfArithmeticSlices(p.A))
	}
	fmt.Printf("\n\n\n")
}
