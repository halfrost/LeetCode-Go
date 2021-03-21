package leetcode

import (
	"fmt"
	"testing"
)

type question643 struct {
	para643
	ans643
}

// para 是参数
// one 代表第一个参数
type para643 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans643 struct {
	one float64
}

func Test_Problem643(t *testing.T) {

	qs := []question643{

		{
			para643{[]int{1, 12, -5, -6, 50, 3}, 4},
			ans643{12.75},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 643------------------------\n")

	for _, q := range qs {
		_, p := q.ans643, q.para643
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaxAverage(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
