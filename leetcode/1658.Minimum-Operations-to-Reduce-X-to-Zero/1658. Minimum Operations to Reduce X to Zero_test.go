package leetcode

import (
	"fmt"
	"testing"
)

type question1658 struct {
	para1658
	ans1658
}

// para 是参数
// one 代表第一个参数
type para1658 struct {
	nums []int
	x    int
}

// ans 是答案
// one 代表第一个答案
type ans1658 struct {
	one int
}

func Test_Problem1658(t *testing.T) {

	qs := []question1658{

		{
			para1658{[]int{1, 1, 4, 2, 3}, 5},
			ans1658{2},
		},

		{
			para1658{[]int{5, 6, 7, 8, 9}, 4},
			ans1658{-1},
		},

		{
			para1658{[]int{3, 2, 20, 1, 1, 3}, 10},
			ans1658{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1658------------------------\n")

	for _, q := range qs {
		_, p := q.ans1658, q.para1658
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minOperations(p.nums, p.x))
	}
	fmt.Printf("\n\n\n")
}
