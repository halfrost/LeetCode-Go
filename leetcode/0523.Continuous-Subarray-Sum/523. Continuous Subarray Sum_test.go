package leetcode

import (
	"fmt"
	"testing"
)

type question523 struct {
	para523
	ans523
}

// para 是参数
// one 代表第一个参数
type para523 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans523 struct {
	one bool
}

func Test_Problem523(t *testing.T) {

	qs := []question523{

		{
			para523{[]int{23, 2, 4, 6, 7}, 6},
			ans523{true},
		},

		{
			para523{[]int{23, 2, 6, 4, 7}, 6},
			ans523{true},
		},

		{
			para523{[]int{23, 2, 6, 4, 7}, 13},
			ans523{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 523------------------------\n")

	for _, q := range qs {
		_, p := q.ans523, q.para523
		fmt.Printf("【input】:%v       【output】:%v\n", p, checkSubarraySum(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
