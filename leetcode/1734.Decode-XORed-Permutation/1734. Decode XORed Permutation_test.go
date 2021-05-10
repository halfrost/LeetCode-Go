package leetcode

import (
	"fmt"
	"testing"
)

type question1734 struct {
	para1734
	ans1734
}

// para 是参数
// one 代表第一个参数
type para1734 struct {
	encoded []int
}

// ans 是答案
// one 代表第一个答案
type ans1734 struct {
	one []int
}

func Test_Problem1734(t *testing.T) {

	qs := []question1734{

		{
			para1734{[]int{3, 1}},
			ans1734{[]int{1, 2, 3}},
		},

		{
			para1734{[]int{6, 5, 4, 6}},
			ans1734{[]int{2, 4, 1, 5, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1734------------------------\n")

	for _, q := range qs {
		_, p := q.ans1734, q.para1734
		fmt.Printf("【input】:%v       【output】:%v\n", p, decode(p.encoded))
	}
	fmt.Printf("\n\n\n")
}
