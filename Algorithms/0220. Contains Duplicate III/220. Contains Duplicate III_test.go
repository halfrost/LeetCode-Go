package leetcode

import (
	"fmt"
	"testing"
)

type question220 struct {
	para220
	ans220
}

// para 是参数
// one 代表第一个参数
type para220 struct {
	one []int
	k   int
	t   int
}

// ans 是答案
// one 代表第一个答案
type ans220 struct {
	one bool
}

func Test_Problem220(t *testing.T) {

	qs := []question220{

		question220{
			para220{[]int{7, 1, 3}, 2, 3},
			ans220{true},
		},

		question220{
			para220{[]int{-1, -1}, 1, -1},
			ans220{false},
		},

		question220{
			para220{[]int{1, 2, 3, 1}, 3, 0},
			ans220{true},
		},

		question220{
			para220{[]int{1, 0, 1, 1}, 1, 2},
			ans220{true},
		},

		question220{
			para220{[]int{1, 5, 9, 1, 5, 9}, 2, 3},
			ans220{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 220------------------------\n")

	for _, q := range qs {
		_, p := q.ans220, q.para220
		fmt.Printf("【input】:%v       【output】:%v\n", p, containsNearbyAlmostDuplicate(p.one, p.k, p.t))
	}
	fmt.Printf("\n\n\n")
}
