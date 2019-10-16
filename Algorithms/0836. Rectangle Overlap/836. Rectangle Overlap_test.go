package leetcode

import (
	"fmt"
	"testing"
)

type question836 struct {
	para836
	ans836
}

// para 是参数
// one 代表第一个参数
type para836 struct {
	rec1 []int
	rec2 []int
}

// ans 是答案
// one 代表第一个答案
type ans836 struct {
	one bool
}

func Test_Problem836(t *testing.T) {

	qs := []question836{

		question836{
			para836{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}},
			ans836{true},
		},

		question836{
			para836{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}},
			ans836{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 836------------------------\n")

	for _, q := range qs {
		_, p := q.ans836, q.para836
		fmt.Printf("【input】:%v       【output】:%v\n", p, isRectangleOverlap(p.rec1, p.rec2))
	}
	fmt.Printf("\n\n\n")
}
