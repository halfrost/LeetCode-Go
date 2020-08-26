package leetcode

import (
	"fmt"
	"testing"
)

type question697 struct {
	para697
	ans697
}

// para 是参数
// one 代表第一个参数
type para697 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans697 struct {
	one int
}

func Test_Problem697(t *testing.T) {

	qs := []question697{

		{
			para697{[]int{1, 2, 2, 3, 1}},
			ans697{2},
		},

		{
			para697{[]int{1, 2, 2, 3, 1, 4, 2}},
			ans697{6},
		},

		{
			para697{[]int{}},
			ans697{0},
		},

		{
			para697{[]int{1, 1, 1, 1, 1}},
			ans697{5},
		},

		{
			para697{[]int{1, 2, 2, 3, 1, 4, 2, 1, 2, 2, 3, 1, 4, 2}},
			ans697{13},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 697------------------------\n")

	for _, q := range qs {
		_, p := q.ans697, q.para697
		fmt.Printf("【input】:%v       【output】:%v\n", p, findShortestSubArray(p.one))
	}
	fmt.Printf("\n\n\n")
}
