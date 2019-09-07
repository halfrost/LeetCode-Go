package leetcode

import (
	"fmt"
	"testing"
)

type question1049 struct {
	para1049
	ans1049
}

// para 是参数
// one 代表第一个参数
type para1049 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1049 struct {
	one int
}

func Test_Problem1049(t *testing.T) {

	qs := []question1049{

		question1049{
			para1049{[]int{2, 7, 4, 1, 8, 1}},
			ans1049{1},
		},

		question1049{
			para1049{[]int{21, 26, 31, 33, 40}},
			ans1049{5},
		},

		question1049{
			para1049{[]int{1, 2}},
			ans1049{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1049------------------------\n")

	for _, q := range qs {
		_, p := q.ans1049, q.para1049
		fmt.Printf("【input】:%v       【output】:%v\n", p, lastStoneWeightII(p.one))
	}
	fmt.Printf("\n\n\n")
}
