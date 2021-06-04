package leetcode

import (
	"fmt"
	"testing"
)

type question525 struct {
	para525
	ans525
}

// para 是参数
// one 代表第一个参数
type para525 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans525 struct {
	one int
}

func Test_Problem525(t *testing.T) {

	qs := []question525{

		{
			para525{[]int{0, 1}},
			ans525{2},
		},

		{
			para525{[]int{0, 1, 0}},
			ans525{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 525------------------------\n")

	for _, q := range qs {
		_, p := q.ans525, q.para525
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaxLength(p.nums))
	}
	fmt.Printf("\n\n\n")
}
