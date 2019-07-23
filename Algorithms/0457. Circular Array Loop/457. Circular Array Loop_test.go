package leetcode

import (
	"fmt"
	"testing"
)

type question457 struct {
	para457
	ans457
}

// para 是参数
// one 代表第一个参数
type para457 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans457 struct {
	one bool
}

func Test_Problem457(t *testing.T) {

	qs := []question457{

		question457{
			para457{[]int{-1}},
			ans457{false},
		},

		question457{
			para457{[]int{3, 1, 2}},
			ans457{true},
		},

		question457{
			para457{[]int{-8, -1, 1, 7, 2}},
			ans457{false},
		},

		question457{
			para457{[]int{-1, -2, -3, -4, -5}},
			ans457{false},
		},

		question457{
			para457{[]int{}},
			ans457{false},
		},

		question457{
			para457{[]int{2, -1, 1, 2, 2}},
			ans457{true},
		},

		question457{
			para457{[]int{-1, 2}},
			ans457{false},
		},
		question457{
			para457{[]int{-2, 1, -1, -2, -2}},
			ans457{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 457------------------------\n")

	for _, q := range qs {
		_, p := q.ans457, q.para457
		fmt.Printf("【input】:%v       【output】:%v\n", p, circularArrayLoop(p.one))
	}
	fmt.Printf("\n\n\n")
}
