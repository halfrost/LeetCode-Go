package leetcode

import (
	"fmt"
	"testing"
)

type question368 struct {
	para368
	ans368
}

// para 是参数
// one 代表第一个参数
type para368 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans368 struct {
	one []int
}

func Test_Problem368(t *testing.T) {

	qs := []question368{

		{
			para368{[]int{1, 2, 3}},
			ans368{[]int{1, 2}},
		},

		{
			para368{[]int{1, 2, 4, 8}},
			ans368{[]int{1, 2, 4, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 368------------------------\n")

	for _, q := range qs {
		_, p := q.ans368, q.para368
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestDivisibleSubset(p.one))
	}
	fmt.Printf("\n\n\n")
}
