package leetcode

import (
	"fmt"
	"testing"
)

type question1030 struct {
	para1030
	ans1030
}

// para 是参数
// one 代表第一个参数
type para1030 struct {
	R  int
	C  int
	r0 int
	c0 int
}

// ans 是答案
// one 代表第一个答案
type ans1030 struct {
	one [][]int
}

func Test_Problem1030(t *testing.T) {

	qs := []question1030{

		question1030{
			para1030{1, 2, 0, 0},
			ans1030{[][]int{[]int{0, 0}, []int{0, 1}}},
		},

		question1030{
			para1030{2, 2, 0, 1},
			ans1030{[][]int{[]int{0, 1}, []int{0, 0}, []int{1, 1}, []int{1, 0}}},
		},

		question1030{
			para1030{2, 3, 1, 2},
			ans1030{[][]int{[]int{1, 2}, []int{0, 2}, []int{1, 1}, []int{0, 1}, []int{1, 0}, []int{0, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1030------------------------\n")

	for _, q := range qs {
		_, p := q.ans1030, q.para1030
		fmt.Printf("【input】:%v       【output】:%v\n", p, allCellsDistOrder(p.R, p.C, p.r0, p.c0))
	}
	fmt.Printf("\n\n\n")
}
