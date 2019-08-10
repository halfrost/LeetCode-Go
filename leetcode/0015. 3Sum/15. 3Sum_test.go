package leetcode

import (
	"fmt"
	"testing"
)

type question15 struct {
	para15
	ans15
}

// para 是参数
// one 代表第一个参数
type para15 struct {
	a []int
}

// ans 是答案
// one 代表第一个答案
type ans15 struct {
	one [][]int
}

func Test_Problem15(t *testing.T) {

	qs := []question15{

		question15{
			para15{[]int{0, 0, 0}},
			ans15{[][]int{[]int{0, 0, 0}}},
		},

		question15{
			para15{[]int{-1, 0, 1, 2, -1, -4}},
			ans15{[][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
		},

		question15{
			para15{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			ans15{[][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
		},

		question15{
			para15{[]int{5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5}},
			ans15{[][]int{[]int{-10, 2, 8}, []int{-10, 3, 7}, []int{-10, 4, 6}, []int{-10, 5, 5}, []int{-8, 2, 6}, []int{-8, 3, 5}, []int{-8, 4, 4}, []int{-7, -1, 8},
				[]int{-7, 2, 5}, []int{-7, 3, 4}, []int{-6, -2, 8}, []int{-6, -1, 7}, []int{-6, 2, 4}, []int{-5, -3, 8}, []int{-5, -2, 7}, []int{-5, -1, 6}, []int{-5, 2, 3},
				[]int{-4, -3, 7}, []int{-4, -2, 6}, []int{-4, -1, 5}, []int{-4, 2, 2}, []int{-3, -3, 6}, []int{-3, -2, 5}, []int{-3, -1, 4}, []int{-2, -1, 3}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 15------------------------\n")

	for _, q := range qs {
		_, p := q.ans15, q.para15
		fmt.Printf("【input】:%v       【output】:%v\n", p, threeSum(p.a))
	}
	fmt.Printf("\n\n\n")
}
