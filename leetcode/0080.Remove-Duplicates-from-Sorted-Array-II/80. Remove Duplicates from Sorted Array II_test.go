package leetcode

import (
	"fmt"
	"testing"
)

type question80 struct {
	para80
	ans80
}

// para 是参数
// one 代表第一个参数
type para80 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans80 struct {
	one int
}

func Test_Problem80(t *testing.T) {

	qs := []question80{

		{
			para80{[]int{1, 1, 2}},
			ans80{3},
		},

		{
			para80{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans80{8},
		},

		{
			para80{[]int{0, 0, 0, 0, 0}},
			ans80{2},
		},

		{
			para80{[]int{1}},
			ans80{1},
		},

		{
			para80{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			ans80{7},
		},

		{
			para80{[]int{1, 1, 1, 1, 2, 2, 3}},
			ans80{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 80------------------------\n")

	for _, q := range qs {
		_, p := q.ans80, q.para80
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, removeDuplicates(p.one))
	}
	fmt.Printf("\n\n\n")
}
