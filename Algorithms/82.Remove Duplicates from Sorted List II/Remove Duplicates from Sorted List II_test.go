package leetcode

import (
	"fmt"
	"testing"
)

type question82 struct {
	para82
	ans82
}

// para 是参数
// one 代表第一个参数
type para82 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans82 struct {
	one []int
}

func Test_Problem82(t *testing.T) {

	qs := []question82{

		question82{
			para82{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			ans82{[]int{3}},
		},

		question82{
			para82{[]int{1, 1, 1, 1, 1, 1}},
			ans82{[]int{}},
		},

		question82{
			para82{[]int{1, 1, 1, 2, 3}},
			ans82{[]int{2, 3}},
		},

		question82{
			para82{[]int{1}},
			ans82{[]int{1}},
		},

		question82{
			para82{[]int{}},
			ans82{[]int{}},
		},

		question82{
			para82{[]int{1, 2, 2, 2, 2}},
			ans82{[]int{1}},
		},

		question82{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 5, 6}},
			ans82{[]int{2, 4, 6}},
		},

		question82{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 6}},
			ans82{[]int{2, 4, 5, 6}},
		},

		question82{
			para82{[]int{0, 1, 2, 2, 3, 4}},
			ans82{[]int{0, 1, 2, 2, 3, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 82------------------------\n")

	for _, q := range qs {
		_, p := q.ans82, q.para82
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(deleteDuplicates1(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
