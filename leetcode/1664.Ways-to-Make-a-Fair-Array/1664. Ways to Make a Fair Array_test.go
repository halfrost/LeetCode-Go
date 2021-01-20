package leetcode

import (
	"fmt"
	"testing"
)

type question1664 struct {
	para1664
	ans1664
}

// para 是参数
// one 代表第一个参数
type para1664 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1664 struct {
	one int
}

func Test_Problem1664(t *testing.T) {

	qs := []question1664{

		{
			para1664{[]int{6, 1, 7, 4, 1}},
			ans1664{0},
		},

		{
			para1664{[]int{2, 1, 6, 4}},
			ans1664{1},
		},

		{
			para1664{[]int{1, 1, 1}},
			ans1664{3},
		},

		{
			para1664{[]int{1, 2, 3}},
			ans1664{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1664------------------------\n")

	for _, q := range qs {
		_, p := q.ans1664, q.para1664
		fmt.Printf("【input】:%v      【output】:%v      \n", p, waysToMakeFair(p.nums))
	}
	fmt.Printf("\n\n\n")
}
