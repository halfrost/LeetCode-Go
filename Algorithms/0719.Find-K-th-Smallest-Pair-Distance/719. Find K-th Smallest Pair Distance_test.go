package leetcode

import (
	"fmt"
	"testing"
)

type question719 struct {
	para719
	ans719
}

// para 是参数
// one 代表第一个参数
type para719 struct {
	num []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans719 struct {
	one int
}

func Test_Problem719(t *testing.T) {

	qs := []question719{

		question719{
			para719{[]int{1, 3, 1}, 1},
			ans719{0},
		},

		question719{
			para719{[]int{1, 1, 1}, 2},
			ans719{0},
		},

		question719{
			para719{[]int{1, 6, 1}, 3},
			ans719{5},
		},

		question719{
			para719{[]int{62, 100, 4}, 2},
			ans719{58},
		},

		question719{
			para719{[]int{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}, 18},
			ans719{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 719------------------------\n")

	for _, q := range qs {
		_, p := q.ans719, q.para719
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestDistancePair(p.num, p.k))
	}
	fmt.Printf("\n\n\n")
}
