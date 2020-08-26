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

		{
			para218{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}},
			ans218{[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 218------------------------\n")

	for _, q := range qs {
		_, p := q.ans218, q.para218
		fmt.Printf("【input】:%v       【output】:%v\n", p, getSkyline(p.one))
	}
	fmt.Printf("\n\n\n")
}
