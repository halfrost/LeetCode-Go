package leetcode

import (
	"fmt"
	"testing"
)

type question75 struct {
	para75
	ans75
}

// para 是参数
// one 代表第一个参数
type para75 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans75 struct {
	one []int
}

func Test_Problem75(t *testing.T) {

	qs := []question75{

		question75{
			para75{[]int{}},
			ans75{[]int{}},
		},

		question75{
			para75{[]int{1}},
			ans75{[]int{1}},
		},

		question75{
			para75{[]int{2, 0, 2, 1, 1, 0}},
			ans75{[]int{0, 0, 1, 1, 2, 2}},
		},

		question75{
			para75{[]int{2, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 0, 1, 2, 2, 2, 0, 1, 1}},
			ans75{[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 75------------------------\n")

	for _, q := range qs {
		_, p := q.ans75, q.para75
		fmt.Printf("【input】:%v      ", p)
		sortColors(p.one)
		fmt.Printf("【output】:%v      \n", p)
	}
	fmt.Printf("\n\n\n")
}
