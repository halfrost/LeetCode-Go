package leetcode

import (
	"fmt"
	"testing"
)

type question1984 struct {
	para1984
	ans1984
}

// para 是参数
type para1984 struct {
	nums []int
	k    int
}

// ans 是答案
type ans1984 struct {
	ans int
}

func Test_Problem1984(t *testing.T) {

	qs := []question1984{

		{
			para1984{[]int{90}, 1},
			ans1984{0},
		},

		{
			para1984{[]int{9, 4, 1, 7}, 2},
			ans1984{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1984------------------------\n")

	for _, q := range qs {
		_, p := q.ans1984, q.para1984
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", minimumDifference(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
