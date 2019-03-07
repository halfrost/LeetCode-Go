package leetcode

import (
	"fmt"
	"testing"
)

type question61 struct {
	para61
	ans61
}

// para 是参数
// one 代表第一个参数
type para61 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans61 struct {
	one []int
}

func Test_Problem61(t *testing.T) {

	qs := []question61{

		question61{
			para61{[]int{1, 2, 3, 4, 5}, 2},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		question61{
			para61{[]int{1, 2, 3, 4, 5}, 3},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		question61{
			para61{[]int{0, 1, 2}, 4},
			ans61{[]int{2, 0, 1}},
		},

		question61{
			para61{[]int{1, 1, 1, 2}, 3},
			ans61{[]int{1, 1, 2, 1}},
		},

		question61{
			para61{[]int{1}, 10},
			ans61{[]int{1}},
		},

		question61{
			para61{[]int{}, 100},
			ans61{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 61------------------------\n")

	for _, q := range qs {
		_, p := q.ans61, q.para61
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(rotateRight(S2l(p.one), p.k)))
	}
	fmt.Printf("\n\n\n")
}
