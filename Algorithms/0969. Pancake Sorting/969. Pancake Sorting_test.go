package leetcode

import (
	"fmt"
	"testing"
)

type question969 struct {
	para969
	ans969
}

// para 是参数
// one 代表第一个参数
type para969 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans969 struct {
	one []int
}

func Test_Problem969(t *testing.T) {

	qs := []question969{

		question969{
			para969{[]int{}},
			ans969{[]int{}},
		},

		question969{
			para969{[]int{1}},
			ans969{[]int{1}},
		},

		question969{
			para969{[]int{3, 2, 4, 1}},
			ans969{[]int{3, 4, 2, 3, 1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 969------------------------\n")

	for _, q := range qs {
		_, p := q.ans969, q.para969
		fmt.Printf("【input】:%v       【output】:%v\n", p, pancakeSort(p.one))
	}
	fmt.Printf("\n\n\n")
}
