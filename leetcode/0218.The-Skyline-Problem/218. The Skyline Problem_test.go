package leetcode

import (
	"fmt"
	"testing"
)

type question218 struct {
	para218
	ans218
}

// para 是参数
// one 代表第一个参数
type para218 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans218 struct {
	one [][]int
}

func Test_Problem218(t *testing.T) {

	qs := []question218{

		question218{
			para218{[][]int{[]int{2, 9, 10}, []int{3, 7, 15}, []int{5, 12, 12}, []int{15, 20, 10}, []int{19, 24, 8}}},
			ans218{[][]int{[]int{2, 10}, []int{3, 15}, []int{7, 12}, []int{12, 0}, []int{15, 10}, []int{20, 8}, []int{24, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 218------------------------\n")

	for _, q := range qs {
		_, p := q.ans218, q.para218
		fmt.Printf("【input】:%v       【output】:%v\n", p, getSkyline(p.one))
	}
	fmt.Printf("\n\n\n")
}
