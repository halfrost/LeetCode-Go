package leetcode

import (
	"fmt"
	"testing"
)

type question1877 struct {
	para1877
	ans1877
}

// para 是参数
// one 代表第一个参数
type para1877 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1877 struct {
	one int
}

func Test_Problem1877(t *testing.T) {

	qs := []question1877{

		{
			para1877{[]int{2, 2, 1, 2, 1}},
			ans1877{3},
		},

		{
			para1877{[]int{100, 1, 1000}},
			ans1877{1001},
		},

		{
			para1877{[]int{1, 2, 3, 4, 5}},
			ans1877{6},
		},

		{
			para1877{[]int{3, 5, 2, 3}},
			ans1877{7},
		},

		{
			para1877{[]int{3, 5, 4, 2, 4, 6}},
			ans1877{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1877------------------------\n")

	for _, q := range qs {
		_, p := q.ans1877, q.para1877
		fmt.Printf("【input】:%v       【output】:%v\n", p, minPairSum(p.nums))
	}
	fmt.Printf("\n\n\n")
}
