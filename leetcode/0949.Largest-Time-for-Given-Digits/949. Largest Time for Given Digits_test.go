package leetcode

import (
	"fmt"
	"testing"
)

type question949 struct {
	para949
	ans949
}

// para 是参数
// one 代表第一个参数
type para949 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans949 struct {
	one string
}

func Test_Problem949(t *testing.T) {

	qs := []question949{
		{
			para949{[]int{1, 2, 3, 4}},
			ans949{"23:41"},
		},

		{
			para949{[]int{5, 5, 5, 5}},
			ans949{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 949------------------------\n")

	for _, q := range qs {
		_, p := q.ans949, q.para949
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestTimeFromDigits(p.one))
	}
	fmt.Printf("\n\n\n")
}
