package leetcode

import (
	"fmt"
	"testing"
)

type question765 struct {
	para765
	ans765
}

// para 是参数
// one 代表第一个参数
type para765 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans765 struct {
	one int
}

func Test_Problem765(t *testing.T) {

	qs := []question765{

		question765{
			para765{[]int{0, 2, 1, 3}},
			ans765{1},
		},

		question765{
			para765{[]int{3, 2, 0, 1}},
			ans765{0},
		},

		question765{
			para765{[]int{3, 1, 4, 0, 2, 5}},
			ans765{2},
		},

		question765{
			para765{[]int{10, 7, 4, 2, 3, 0, 9, 11, 1, 5, 6, 8}},
			ans765{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 765------------------------\n")

	for _, q := range qs {
		_, p := q.ans765, q.para765
		fmt.Printf("【input】:%v       【output】:%v\n", p, minSwapsCouples(p.one))
	}
	fmt.Printf("\n\n\n")
}
