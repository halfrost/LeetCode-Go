package leetcode

import (
	"fmt"
	"testing"
)

type question147 struct {
	para147
	ans147
}

// para 是参数
// one 代表第一个参数
type para147 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans147 struct {
	one []int
}

func Test_Problem147(t *testing.T) {

	qs := []question147{

		question147{
			para147{[]int{4, 2, 1, 3}},
			ans147{[]int{1, 2, 3, 4}},
		},
		question147{
			para147{[]int{1}},
			ans147{[]int{1}},
		},

		question147{
			para147{[]int{-1, 5, 3, 4, 0}},
			ans147{[]int{-1, 0, 3, 4, 5}},
		},

		question147{
			para147{[]int{}},
			ans147{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 147------------------------\n")

	for _, q := range qs {
		_, p := q.ans147, q.para147
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(insertionSortList(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
