package leetcode

import (
	"fmt"
	"testing"
)

type question875 struct {
	para875
	ans875
}

// para 是参数
// one 代表第一个参数
type para875 struct {
	piles []int
	H     int
}

// ans 是答案
// one 代表第一个答案
type ans875 struct {
	one int
}

func Test_Problem875(t *testing.T) {

	qs := []question875{

		question875{
			para875{[]int{3, 6, 7, 11}, 8},
			ans875{4},
		},

		question875{
			para875{[]int{30, 11, 23, 4, 20}, 5},
			ans875{30},
		},

		question875{
			para875{[]int{30, 11, 23, 4, 20}, 6},
			ans875{23},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 875------------------------\n")

	for _, q := range qs {
		_, p := q.ans875, q.para875
		fmt.Printf("【input】:%v       【output】:%v\n", p, minEatingSpeed(p.piles, p.H))
	}
	fmt.Printf("\n\n\n")
}
