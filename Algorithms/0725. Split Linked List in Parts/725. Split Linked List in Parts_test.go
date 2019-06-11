package leetcode

import (
	"fmt"
	"testing"
)

type question725 struct {
	para725
	ans725
}

// para 是参数
// one 代表第一个参数
type para725 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans725 struct {
	one []int
}

func Test_Problem725(t *testing.T) {

	qs := []question725{

		question725{
			para725{[]int{1, 2, 3, 4, 5}, 7},
			ans725{[]int{1, 2, 3, 4, 5, 0, 0}},
		},

		question725{
			para725{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3},
			ans725{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},

		question725{
			para725{[]int{1, 1, 1, 1, 1}, 1},
			ans725{[]int{1, 1, 1, 1, 1}},
		},

		question725{
			para725{[]int{}, 3},
			ans725{[]int{}},
		},

		question725{
			para725{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 0},
			ans725{[]int{1, 2, 3, 2, 3, 2, 3, 2}},
		},

		question725{
			para725{[]int{1, 2, 3, 4, 5}, 5},
			ans725{[]int{1, 2, 3, 4}},
		},

		question725{
			para725{[]int{}, 5},
			ans725{[]int{}},
		},

		question725{
			para725{[]int{1, 2, 3, 4, 5}, 10},
			ans725{[]int{1, 2, 3, 4, 5}},
		},

		question725{
			para725{[]int{1}, 1},
			ans725{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 725------------------------\n")

	for _, q := range qs {
		_, p := q.ans725, q.para725
		res := splitListToParts(S2l(p.one), p.n)
		for _, value := range res {
			fmt.Printf("【input】:%v    length:%v   【output】:%v\n", p, len(res), L2s(value))
		}
		fmt.Printf("\n\n\n")
	}
	fmt.Printf("\n\n\n")
}
