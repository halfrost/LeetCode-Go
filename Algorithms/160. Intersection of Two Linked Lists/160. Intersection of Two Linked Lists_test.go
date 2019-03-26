package leetcode

import (
	"fmt"
	"testing"
)

type question160 struct {
	para160
	ans160
}

// para 是参数
// one 代表第一个参数
type para160 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans160 struct {
	one []int
}

func Test_Problem160(t *testing.T) {

	qs := []question160{

		question160{
			para160{[]int{}, []int{}},
			ans160{[]int{}},
		},

		question160{
			para160{[]int{3}, []int{1, 2, 3}},
			ans160{[]int{3}},
		},

		question160{
			para160{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans160{[]int{1, 2, 3, 4}},
		},

		question160{
			para160{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}},
			ans160{[]int{8, 4, 5}},
		},

		question160{
			para160{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans160{[]int{}},
		},

		question160{
			para160{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}},
			ans160{[]int{2, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 160------------------------\n")

	for _, q := range qs {
		_, p := q.ans160, q.para160
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(getIntersectionNode(S2l(p.one), S2l(p.another))))
	}
	fmt.Printf("\n\n\n")
}
