package leetcode

import (
	"fmt"
	"testing"
)

type question1652 struct {
	para1652
	ans1652
}

// para 是参数
// one 代表第一个参数
type para1652 struct {
	code []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1652 struct {
	one []int
}

func Test_Problem1652(t *testing.T) {

	qs := []question1652{

		{
			para1652{[]int{5, 7, 1, 4}, 3},
			ans1652{[]int{12, 10, 16, 13}},
		},

		{
			para1652{[]int{1, 2, 3, 4}, 0},
			ans1652{[]int{0, 0, 0, 0}},
		},

		{
			para1652{[]int{2, 4, 9, 3}, -2},
			ans1652{[]int{12, 5, 6, 13}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1652------------------------\n")

	for _, q := range qs {
		_, p := q.ans1652, q.para1652
		fmt.Printf("【input】:%v      【output】:%v      \n", p, decrypt(p.code, p.k))
	}
	fmt.Printf("\n\n\n")
}
