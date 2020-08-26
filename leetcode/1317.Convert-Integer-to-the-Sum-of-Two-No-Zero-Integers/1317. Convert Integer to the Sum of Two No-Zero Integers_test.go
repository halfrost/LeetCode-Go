package leetcode

import (
	"fmt"
	"testing"
)

type question1317 struct {
	para1317
	ans1317
}

// para 是参数
// one 代表第一个参数
type para1317 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1317 struct {
	one []int
}

func Test_Problem1317(t *testing.T) {

	qs := []question1317{

		{
			para1317{5},
			ans1317{[]int{1, 4}},
		},

		{
			para1317{0},
			ans1317{[]int{}},
		},

		{
			para1317{3},
			ans1317{[]int{1, 2}},
		},

		{
			para1317{1},
			ans1317{[]int{}},
		},

		{
			para1317{2},
			ans1317{[]int{1, 1}},
		},

		{
			para1317{11},
			ans1317{[]int{2, 9}},
		},

		{
			para1317{10000},
			ans1317{[]int{1, 9999}},
		},

		{
			para1317{69},
			ans1317{[]int{1, 68}},
		},

		{
			para1317{1010},
			ans1317{[]int{11, 999}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1317------------------------\n")

	for _, q := range qs {
		_, p := q.ans1317, q.para1317
		fmt.Printf("【input】:%v       【output】:%v\n", p, getNoZeroIntegers(p.one))
	}
	fmt.Printf("\n\n\n")
}
