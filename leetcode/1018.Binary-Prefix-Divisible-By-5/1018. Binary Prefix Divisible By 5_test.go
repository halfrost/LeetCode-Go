package leetcode

import (
	"fmt"
	"testing"
)

type question1018 struct {
	para1018
	ans1018
}

// para 是参数
// one 代表第一个参数
type para1018 struct {
	a []int
}

// ans 是答案
// one 代表第一个答案
type ans1018 struct {
	one []bool
}

func Test_Problem1018(t *testing.T) {

	qs := []question1018{

		{
			para1018{[]int{0, 1, 1}},
			ans1018{[]bool{true, false, false}},
		},

		{
			para1018{[]int{1, 1, 1}},
			ans1018{[]bool{false, false, false}},
		},

		{
			para1018{[]int{0, 1, 1, 1, 1, 1}},
			ans1018{[]bool{true, false, false, false, true, false}},
		},

		{
			para1018{[]int{1, 1, 1, 0, 1}},
			ans1018{[]bool{false, false, false, false, false}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1018------------------------\n")

	for _, q := range qs {
		_, p := q.ans1018, q.para1018
		fmt.Printf("【input】:%v       【output】:%v\n", p, prefixesDivBy5(p.a))
	}
	fmt.Printf("\n\n\n")
}
