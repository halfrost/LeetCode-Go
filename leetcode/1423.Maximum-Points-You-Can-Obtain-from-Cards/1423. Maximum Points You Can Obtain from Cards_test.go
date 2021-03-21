package leetcode

import (
	"fmt"
	"testing"
)

type question1423 struct {
	para1423
	ans1423
}

// para 是参数
// one 代表第一个参数
type para1423 struct {
	cardPoints []int
	k          int
}

// ans 是答案
// one 代表第一个答案
type ans1423 struct {
	one int
}

func Test_Problem1423(t *testing.T) {

	qs := []question1423{

		{
			para1423{[]int{1, 2, 3, 4, 5, 6, 1}, 3},
			ans1423{12},
		},

		{
			para1423{[]int{2, 2, 2}, 2},
			ans1423{4},
		},

		{
			para1423{[]int{9, 7, 7, 9, 7, 7, 9}, 7},
			ans1423{55},
		},

		{
			para1423{[]int{1, 1000, 1}, 1},
			ans1423{1},
		},

		{
			para1423{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3},
			ans1423{202},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1423------------------------\n")

	for _, q := range qs {
		_, p := q.ans1423, q.para1423
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxScore(p.cardPoints, p.k))
	}
	fmt.Printf("\n\n\n")
}
