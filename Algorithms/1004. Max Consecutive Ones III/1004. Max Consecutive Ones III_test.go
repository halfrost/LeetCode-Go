package leetcode

import (
	"fmt"
	"testing"
)

type question1004 struct {
	para1004
	ans1004
}

// para 是参数
// one 代表第一个参数
type para1004 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1004 struct {
	one int
}

func Test_Problem1004(t *testing.T) {

	qs := []question1004{

		question1004{
			para1004{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
			ans1004{6},
		},

		question1004{
			para1004{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
			ans1004{10},
		},

		question1004{
			para1004{[]int{0, 0, 0, 1}, 4},
			ans1004{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1004------------------------\n")

	for _, q := range qs {
		_, p := q.ans1004, q.para1004
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestOnes(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
