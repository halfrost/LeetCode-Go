package leetcode

import (
	"fmt"
	"testing"
)

type question2167 struct {
	para2167
	ans2167
}

// para 是参数
// one 代表第一个参数
type para2167 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans2167 struct {
	one int
}

func Test_Problem2167(t *testing.T) {

	qs := []question2167{

		{
			para2167{"1100101"},
			ans2167{5},
		},

		{
			para2167{"0010"},
			ans2167{2},
		},

		{
			para2167{"1100111101"},
			ans2167{8},
		},

		{
			para2167{"1001010101"},
			ans2167{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2167------------------------\n")

	for _, q := range qs {
		_, p := q.ans2167, q.para2167
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumTime(p.s))
	}
	fmt.Printf("\n\n\n")
}
