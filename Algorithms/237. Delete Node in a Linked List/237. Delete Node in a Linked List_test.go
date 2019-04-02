package leetcode

import (
	"fmt"
	"testing"
)

type question237 struct {
	para237
	ans237
}

// para 是参数
// one 代表第一个参数
type para237 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans237 struct {
	one []int
}

func Test_Problem237(t *testing.T) {

	qs := []question237{

		question237{
			para237{[]int{1, 2, 3, 4, 5}, 1},
			ans237{[]int{2, 3, 4, 5}},
		},

		question237{
			para237{[]int{1, 2, 3, 4, 5}, 2},
			ans237{[]int{1, 3, 4, 5}},
		},

		question237{
			para237{[]int{1, 1, 1, 1, 1}, 1},
			ans237{[]int{}},
		},

		question237{
			para237{[]int{1, 2, 3, 4, 5}, 5},
			ans237{[]int{1, 2, 3, 4}},
		},

		question237{
			para237{[]int{}, 5},
			ans237{[]int{}},
		},

		question237{
			para237{[]int{1, 2, 3, 4, 5}, 10},
			ans237{[]int{1, 2, 3, 4, 5}},
		},

		question237{
			para237{[]int{1}, 1},
			ans237{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 237------------------------\n")

	for _, q := range qs {
		_, p := q.ans237, q.para237
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(removeElements(S2l(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}
