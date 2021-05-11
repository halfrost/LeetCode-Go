package leetcode

import (
	"fmt"
	"testing"
)

type question1482 struct {
	para1482
	ans1482
}

// para 是参数
// one 代表第一个参数
type para1482 struct {
	bloomDay []int
	m        int
	k        int
}

// ans 是答案
// one 代表第一个答案
type ans1482 struct {
	one int
}

func Test_Problem1482(t *testing.T) {

	qs := []question1482{

		{
			para1482{[]int{1, 10, 3, 10, 2}, 3, 1},
			ans1482{3},
		},

		{
			para1482{[]int{1, 10, 3, 10, 2}, 3, 2},
			ans1482{-1},
		},

		{
			para1482{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3},
			ans1482{12},
		},

		{
			para1482{[]int{1000000000, 1000000000}, 1, 1},
			ans1482{1000000000},
		},

		{
			para1482{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2},
			ans1482{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1482------------------------\n")

	for _, q := range qs {
		_, p := q.ans1482, q.para1482
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minDays(p.bloomDay, p.m, p.k))
	}
	fmt.Printf("\n\n\n")
}
