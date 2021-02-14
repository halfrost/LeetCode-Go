package leetcode

import (
	"fmt"
	"testing"
)

type question1636 struct {
	para1636
	ans1636
}

// para 是参数
// one 代表第一个参数
type para1636 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1636 struct {
	one []int
}

func Test_Problem1636(t *testing.T) {

	qs := []question1636{

		{
			para1636{[]int{1, 1, 2, 2, 2, 3}},
			ans1636{[]int{3, 1, 1, 2, 2, 2}},
		},

		{
			para1636{[]int{2, 3, 1, 3, 2}},
			ans1636{[]int{1, 3, 3, 2, 2}},
		},

		{
			para1636{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
			ans1636{[]int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1636------------------------\n")

	for _, q := range qs {
		_, p := q.ans1636, q.para1636
		fmt.Printf("【input】:%v      【output】:%v      \n", p, frequencySort(p.nums))
	}
	fmt.Printf("\n\n\n")
}
