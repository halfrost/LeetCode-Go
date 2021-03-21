package leetcode

import (
	"fmt"
	"testing"
)

type question1608 struct {
	para1608
	ans1608
}

// para 是参数
// one 代表第一个参数
type para1608 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1608 struct {
	one int
}

func Test_Problem1608(t *testing.T) {

	qs := []question1608{

		{
			para1608{[]int{3, 5}},
			ans1608{2},
		},

		{
			para1608{[]int{0, 0}},
			ans1608{-1},
		},

		{
			para1608{[]int{0, 4, 3, 0, 4}},
			ans1608{3},
		},

		{
			para1608{[]int{3, 6, 7, 7, 0}},
			ans1608{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1608------------------------\n")

	for _, q := range qs {
		_, p := q.ans1608, q.para1608
		fmt.Printf("【input】:%v      【output】:%v      \n", p, specialArray(p.nums))
	}
	fmt.Printf("\n\n\n")
}
