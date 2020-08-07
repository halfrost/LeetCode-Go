package leetcode

import (
	"fmt"
	"testing"
)

type question735 struct {
	para735
	ans735
}

// para 是参数
// one 代表第一个参数
type para735 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans735 struct {
	one []int
}

func Test_Problem735(t *testing.T) {

	qs := []question735{

		question735{
			para735{[]int{-1, -1, 1, -2}},
			ans735{[]int{-1, -1, -2}},
		},

		question735{
			para735{[]int{5, 10, -5}},
			ans735{[]int{5, 10}},
		},

		question735{
			para735{[]int{5, 8, -8}},
			ans735{[]int{5}},
		},

		question735{
			para735{[]int{8, -8}},
			ans735{[]int{}},
		},

		question735{
			para735{[]int{10, 2, -5}},
			ans735{[]int{10}},
		},

		question735{
			para735{[]int{-2, -1, 1, 2}},
			ans735{[]int{-2, -1, 1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 735------------------------\n")

	for _, q := range qs {
		_, p := q.ans735, q.para735
		fmt.Printf("【input】:%v       【output】:%v\n", p, asteroidCollision(p.one))
	}
	fmt.Printf("\n\n\n")
}
