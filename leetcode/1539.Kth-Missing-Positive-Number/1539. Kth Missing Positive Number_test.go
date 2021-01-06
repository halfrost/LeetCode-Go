package leetcode

import (
	"fmt"
	"testing"
)

type question1539 struct {
	para1539
	ans1539
}

// para 是参数
// one 代表第一个参数
type para1539 struct {
	arr []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans1539 struct {
	one int
}

func Test_Problem1539(t *testing.T) {

	qs := []question1539{

		{
			para1539{[]int{2, 3, 4, 7, 11}, 5},
			ans1539{9},
		},

		{
			para1539{[]int{1, 2, 3, 4}, 2},
			ans1539{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1539------------------------\n")

	for _, q := range qs {
		_, p := q.ans1539, q.para1539
		fmt.Printf("【input】:%v      【output】:%v      \n", p, findKthPositive(p.arr, p.k))
	}
	fmt.Printf("\n\n\n")
}
