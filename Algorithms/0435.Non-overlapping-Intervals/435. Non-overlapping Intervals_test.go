package leetcode

import (
	"fmt"
	"testing"
)

type question435 struct {
	para435
	ans435
}

// para 是参数
// one 代表第一个参数
type para435 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans435 struct {
	one int
}

func Test_Problem435(t *testing.T) {

	qs := []question435{

		question435{
			para435{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}}},
			ans435{1},
		},

		question435{
			para435{[][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}},
			ans435{2},
		},

		question435{
			para435{[][]int{[]int{1, 2}, []int{2, 3}}},
			ans435{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 435------------------------\n")

	for _, q := range qs {
		_, p := q.ans435, q.para435
		fmt.Printf("【input】:%v       【output】:%v\n", p, eraseOverlapIntervals1(p.one))
	}
	fmt.Printf("\n\n\n")
}
