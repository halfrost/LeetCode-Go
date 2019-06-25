package leetcode

import (
	"fmt"
	"testing"
)

type question47 struct {
	para47
	ans47
}

// para 是参数
// one 代表第一个参数
type para47 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans47 struct {
	one [][]int
}

func Test_Problem47(t *testing.T) {

	qs := []question47{

		question47{
			para47{[]int{1, 1, 2}},
			ans47{[][]int{[]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1}}},
		},

		question47{
			para47{[]int{1, 2, 2}},
			ans47{[][]int{[]int{1, 2, 2}, []int{2, 2, 1}, []int{2, 1, 2}}},
		},

		question47{
			para47{[]int{2, 2, 2}},
			ans47{[][]int{[]int{2, 2, 2}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 47------------------------\n")

	for _, q := range qs {
		_, p := q.ans47, q.para47
		fmt.Printf("【input】:%v       【output】:%v\n", p, permuteUnique(p.s))
	}
	fmt.Printf("\n\n\n")
}
