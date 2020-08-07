package leetcode

import (
	"fmt"
	"testing"
)

type question268 struct {
	para268
	ans268
}

// para 是参数
// one 代表第一个参数
type para268 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans268 struct {
	one int
}

func Test_Problem268(t *testing.T) {

	qs := []question268{

		question268{
			para268{[]int{3, 0, 1}},
			ans268{2},
		},

		question268{
			para268{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			ans268{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 268------------------------\n")

	for _, q := range qs {
		_, p := q.ans268, q.para268
		fmt.Printf("【input】:%v       【output】:%v\n", p, missingNumber(p.s))
	}
	fmt.Printf("\n\n\n")
}
