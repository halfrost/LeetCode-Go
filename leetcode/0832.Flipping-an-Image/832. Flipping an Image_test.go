package leetcode

import (
	"fmt"
	"testing"
)

type question832 struct {
	para832
	ans832
}

// para 是参数
// one 代表第一个参数
type para832 struct {
	A [][]int
}

// ans 是答案
// one 代表第一个答案
type ans832 struct {
	one [][]int
}

func Test_Problem832(t *testing.T) {

	qs := []question832{

		question832{
			para832{[][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}},
			ans832{[][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}},
		},

		question832{
			para832{[][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}}},
			ans832{[][]int{[]int{1, 1, 0, 0}, []int{0, 1, 1, 0}, []int{0, 0, 0, 1}, []int{1, 0, 1, 0}}},
		},

		question832{
			para832{[][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{0, 0, 0}}},
			ans832{[][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{1, 1, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 832------------------------\n")

	for _, q := range qs {
		_, p := q.ans832, q.para832
		fmt.Printf("【input】:%v       【output】:%v\n", p, flipAndInvertImage(p.A))
	}
	fmt.Printf("\n\n\n")
}
