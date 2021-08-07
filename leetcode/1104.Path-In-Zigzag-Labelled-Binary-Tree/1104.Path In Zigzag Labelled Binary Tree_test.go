package leetcode

import (
	"fmt"
	"testing"
)

type question1104 struct {
	para1104
	ans1104
}

// para 是参数
type para1104 struct {
	label int
}

// ans 是答案
type ans1104 struct {
	ans []int
}

func Test_Problem1104(t *testing.T) {

	qs := []question1104{

		{
			para1104{14},
			ans1104{[]int{1, 3, 4, 14}},
		},

		{
			para1104{26},
			ans1104{[]int{1, 2, 6, 10, 26}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1104------------------------\n")

	for _, q := range qs {
		_, p := q.ans1104, q.para1104
		fmt.Printf("【input】:%v      【output】:%v      \n", p, pathInZigZagTree(p.label))
	}
	fmt.Printf("\n\n\n")
}
