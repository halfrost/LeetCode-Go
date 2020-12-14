package leetcode

import (
	"fmt"
	"testing"
)

type question1679 struct {
	para1679
	ans1679
}

// para 是参数
// one 代表第一个参数
type para1679 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1679 struct {
	one int
}

func Test_Problem1679(t *testing.T) {

	qs := []question1679{

		{
			para1679{[]int{1, 2, 3, 4}, 5},
			ans1679{2},
		},

		{
			para1679{[]int{3, 1, 3, 4, 3}, 6},
			ans1679{1},
		},

		{
			para1679{[]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3},
			ans1679{4},
		},

		{
			para1679{[]int{2, 5, 5, 5, 1, 3, 4, 4, 1, 4, 4, 1, 3, 1, 3, 1, 3, 2, 4, 2}, 6},
			ans1679{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1679------------------------\n")

	for _, q := range qs {
		_, p := q.ans1679, q.para1679
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxOperations(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
