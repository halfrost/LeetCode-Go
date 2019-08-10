package leetcode

import (
	"fmt"
	"testing"
)

type question817 struct {
	para817
	ans817
}

// para 是参数
// one 代表第一个参数
type para817 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans817 struct {
	one int
}

func Test_Problem817(t *testing.T) {

	qs := []question817{

		question817{
			para817{[]int{0, 1, 2, 3}, []int{0, 1, 3}},
			ans817{2},
		},

		question817{
			para817{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans817{1},
		},

		question817{
			para817{[]int{0, 1, 2, 3, 4}, []int{0, 3, 1, 4}},
			ans817{2},
		},

		question817{
			para817{[]int{0, 1, 2}, []int{0, 2}},
			ans817{2},
		},

		question817{
			para817{[]int{1, 2, 0, 4, 3}, []int{3, 4, 0, 2, 1}},
			ans817{1},
		},

		question817{
			para817{[]int{0, 2, 4, 3, 1}, []int{3, 2, 4}},
			ans817{1},
		},

		question817{
			para817{[]int{7, 5, 13, 3, 16, 11, 12, 0, 17, 1, 4, 15, 6, 14, 2, 19, 9, 10, 8, 18}, []int{8, 10, 3, 11, 17, 16, 7, 9, 5, 15, 13, 6}},
			ans817{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 817------------------------\n")

	for _, q := range qs {
		_, p := q.ans817, q.para817
		fmt.Printf("【input】:%v       【output】:%v\n", p, numComponents(S2l(p.one), p.another))
	}
	fmt.Printf("\n\n\n")
}
