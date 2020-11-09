package leetcode

import (
	"fmt"
	"testing"
)

type question1649 struct {
	para1649
	ans1649
}

// para 是参数
// one 代表第一个参数
type para1649 struct {
	instructions []int
}

// ans 是答案
// one 代表第一个答案
type ans1649 struct {
	one int
}

func Test_Problem1649(t *testing.T) {

	qs := []question1649{

		{
			para1649{[]int{1, 5, 6, 2}},
			ans1649{1},
		},

		{
			para1649{[]int{1, 2, 3, 6, 5, 4}},
			ans1649{3},
		},

		{
			para1649{[]int{1, 3, 3, 3, 2, 4, 2, 1, 2}},
			ans1649{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1649------------------------\n")

	for _, q := range qs {
		_, p := q.ans1649, q.para1649
		fmt.Printf("【input】:%v      【output】:%v      \n", p, createSortedArray(p.instructions))
	}
	fmt.Printf("\n\n\n")
}
