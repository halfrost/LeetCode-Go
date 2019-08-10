package leetcode

import (
	"fmt"
	"testing"
)

type question90 struct {
	para90
	ans90
}

// para 是参数
// one 代表第一个参数
type para90 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans90 struct {
	one [][]int
}

func Test_Problem90(t *testing.T) {

	qs := []question90{

		question90{
			para90{[]int{}},
			ans90{[][]int{[]int{}}},
		},

		question90{
			para90{[]int{1, 2, 2}},
			ans90{[][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{2, 2}, []int{1, 2, 2}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 90------------------------\n")

	for _, q := range qs {
		_, p := q.ans90, q.para90
		fmt.Printf("【input】:%v       【output】:%v\n", p, subsetsWithDup(p.one))
	}
	fmt.Printf("\n\n\n")
}
