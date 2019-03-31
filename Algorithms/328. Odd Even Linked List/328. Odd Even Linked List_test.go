package leetcode

import (
	"fmt"
	"testing"
)

type question328 struct {
	para328
	ans328
}

// para 是参数
// one 代表第一个参数
type para328 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans328 struct {
	one []int
}

func Test_Problem328(t *testing.T) {

	qs := []question328{

		question328{
			para328{[]int{1, 4, 3, 2, 5, 2}},
			ans328{[]int{1, 3, 5, 4, 2, 2}},
		},

		question328{
			para328{[]int{1, 1, 2, 2, 3, 3, 3}},
			ans328{[]int{1, 2, 3, 3, 1, 2, 3}},
		},

		question328{
			para328{[]int{4, 3, 2, 5, 2}},
			ans328{[]int{4, 2, 2, 3, 5}},
		},

		question328{
			para328{[]int{1, 1, 1, 1, 1, 1}},
			ans328{[]int{1, 1, 1, 1, 1, 1}},
		},

		question328{
			para328{[]int{3, 1}},
			ans328{[]int{3, 1}},
		},

		question328{
			para328{[]int{1, 2, 3, 4, 5}},
			ans328{[]int{1, 3, 5, 2, 4}},
		},

		question328{
			para328{[]int{}},
			ans328{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 328------------------------\n")

	for _, q := range qs {
		_, p := q.ans328, q.para328
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(oddEvenList(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
