package leetcode

import (
	"fmt"
	"testing"
)

type question1300 struct {
	para1300
	ans1300
}

// para 是参数
// one 代表第一个参数
type para1300 struct {
	arr    []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1300 struct {
	one int
}

func Test_Problem1300(t *testing.T) {

	qs := []question1300{

		{
			para1300{[]int{4, 9, 3}, 10},
			ans1300{3},
		},

		{
			para1300{[]int{2, 3, 5}, 10},
			ans1300{5},
		},

		{
			para1300{[]int{60864, 25176, 27249, 21296, 20204}, 56803},
			ans1300{11361},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1300------------------------\n")

	for _, q := range qs {
		_, p := q.ans1300, q.para1300
		fmt.Printf("【input】:%v       【output】:%v\n", p, findBestValue(p.arr, p.target))
	}
	fmt.Printf("\n\n\n")
}
