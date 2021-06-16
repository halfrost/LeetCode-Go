package leetcode

import (
	"fmt"
	"testing"
)

type question473 struct {
	para473
	ans473
}

// para 是参数
// one 代表第一个参数
type para473 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans473 struct {
	one bool
}

func Test_Problem473(t *testing.T) {

	qs := []question473{

		{
			para473{[]int{1, 1, 2, 2, 2}},
			ans473{true},
		},

		{
			para473{[]int{3, 3, 3, 3, 4}},
			ans473{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 473------------------------\n")

	for _, q := range qs {
		_, p := q.ans473, q.para473
		fmt.Printf("【input】:%v       【output】:%v\n", p, makesquare(p.arr))
	}
	fmt.Printf("\n\n\n")
}
