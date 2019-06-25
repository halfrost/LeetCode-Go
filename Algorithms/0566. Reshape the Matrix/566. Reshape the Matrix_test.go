package leetcode

import (
	"fmt"
	"testing"
)

type question566 struct {
	para566
	ans566
}

// para 是参数
// one 代表第一个参数
type para566 struct {
	nums [][]int
	r    int
	c    int
}

// ans 是答案
// one 代表第一个答案
type ans566 struct {
	one [][]int
}

func Test_Problem566(t *testing.T) {

	qs := []question566{

		question566{
			para566{[][]int{[]int{1, 2}, []int{3, 4}}, 1, 4},
			ans566{[][]int{[]int{1, 2, 3, 4}}},
		},

		question566{
			para566{[][]int{[]int{1, 2}, []int{3, 4}}, 2, 4},
			ans566{[][]int{[]int{1, 2}, []int{3, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 566------------------------\n")

	for _, q := range qs {
		_, p := q.ans566, q.para566
		fmt.Printf("【input】:%v       【output】:%v\n", p, matrixReshape(p.nums, p.r, p.c))
	}
	fmt.Printf("\n\n\n")
}
