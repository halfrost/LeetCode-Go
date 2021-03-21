package leetcode

import (
	"fmt"
	"testing"
)

type question1685 struct {
	para1685
	ans1685
}

// para 是参数
// one 代表第一个参数
type para1685 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1685 struct {
	one []int
}

func Test_Problem1685(t *testing.T) {

	qs := []question1685{

		{
			para1685{[]int{2, 3, 5}},
			ans1685{[]int{4, 3, 5}},
		},

		{
			para1685{[]int{1, 4, 6, 8, 10}},
			ans1685{[]int{24, 15, 13, 15, 21}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1685------------------------\n")

	for _, q := range qs {
		_, p := q.ans1685, q.para1685
		fmt.Printf("【input】:%v       【output】:%v\n", p, getSumAbsoluteDifferences(p.nums))
	}
	fmt.Printf("\n\n\n")
}
