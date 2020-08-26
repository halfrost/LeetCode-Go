package leetcode

import (
	"fmt"
	"testing"
)

type question498 struct {
	para498
	ans498
}

// para 是参数
// one 代表第一个参数
type para498 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans498 struct {
	one []int
}

func Test_Problem498(t *testing.T) {

	qs := []question498{

		{
			para498{[][]int{{3}, {2}, {9}}},
			ans498{[]int{3, 2, 9}},
		},

		{
			para498{[][]int{{6, 9, 7}}},
			ans498{[]int{6, 9, 7}},
		},

		{
			para498{[][]int{{3}, {2}}},
			ans498{[]int{3, 2}},
		},

		{
			para498{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans498{[]int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		},

		{
			para498{[][]int{{0}}},
			ans498{[]int{0}},
		},

		{
			para498{[][]int{{}}},
			ans498{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 498------------------------\n")

	for _, q := range qs {
		_, p := q.ans498, q.para498
		fmt.Printf("【input】:%v       【output】:%v\n", p, findDiagonalOrder(p.one))
	}
	fmt.Printf("\n\n\n")
}
