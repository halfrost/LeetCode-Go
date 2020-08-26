package leetcode

import (
	"fmt"
	"testing"
)

type question896 struct {
	para896
	ans896
}

// para 是参数
// one 代表第一个参数
type para896 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans896 struct {
	one bool
}

func Test_Problem896(t *testing.T) {

	qs := []question896{

		{
			para896{[]int{2, 1, 3}},
			ans896{false},
		},

		{
			para896{[]int{1, 2, 2, 3}},
			ans896{true},
		},

		{
			para896{[]int{6, 5, 4, 4}},
			ans896{true},
		},

		{
			para896{[]int{1, 3, 2}},
			ans896{false},
		},

		{
			para896{[]int{1, 2, 4, 5}},
			ans896{true},
		},

		{
			para896{[]int{1, 1, 1}},
			ans896{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 896------------------------\n")

	for _, q := range qs {
		_, p := q.ans896, q.para896
		fmt.Printf("【input】:%v       【output】:%v\n", p, isMonotonic(p.one))
	}
	fmt.Printf("\n\n\n")
}
