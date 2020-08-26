package leetcode

import (
	"fmt"
	"testing"
)

type question1389 struct {
	para1389
	ans1389
}

// para 是参数
// one 代表第一个参数
type para1389 struct {
	nums  []int
	index []int
}

// ans 是答案
// one 代表第一个答案
type ans1389 struct {
	one []int
}

func Test_Problem1389(t *testing.T) {

	qs := []question1389{

		{
			para1389{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}},
			ans1389{[]int{0, 4, 1, 3, 2}},
		},

		{
			para1389{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}},
			ans1389{[]int{0, 1, 2, 3, 4}},
		},

		{
			para1389{[]int{1}, []int{0}},
			ans1389{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1389------------------------\n")

	for _, q := range qs {
		_, p := q.ans1389, q.para1389
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", createTargetArray(p.nums, p.index))
	}
	fmt.Printf("\n\n\n")
}
