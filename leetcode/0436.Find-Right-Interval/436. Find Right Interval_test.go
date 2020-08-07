package leetcode

import (
	"fmt"
	"testing"
)

type question436 struct {
	para436
	ans436
}

// para 是参数
// one 代表第一个参数
type para436 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans436 struct {
	one []int
}

func Test_Problem436(t *testing.T) {

	qs := []question436{

		question436{
			para436{[][]int{[]int{3, 4}, []int{2, 3}, []int{1, 2}}},
			ans436{[]int{-1, 0, 1}},
		},

		question436{
			para436{[][]int{[]int{1, 4}, []int{2, 3}, []int{3, 4}}},
			ans436{[]int{-1, 2, -1}},
		},

		question436{
			para436{[][]int{[]int{1, 2}}},
			ans436{[]int{-1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 436------------------------\n")

	for _, q := range qs {
		_, p := q.ans436, q.para436
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRightInterval(p.one))
	}
	fmt.Printf("\n\n\n")
}
