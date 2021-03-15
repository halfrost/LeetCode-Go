package leetcode

import (
	"fmt"
	"testing"
)

type question1461 struct {
	para1461
	ans1461
}

// para 是参数
// one 代表第一个参数
type para1461 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1461 struct {
	one bool
}

func Test_Problem1461(t *testing.T) {

	qs := []question1461{

		{
			para1461{"00110110", 2},
			ans1461{true},
		},

		{
			para1461{"00110", 2},
			ans1461{true},
		},

		{
			para1461{"0110", 1},
			ans1461{true},
		},

		{
			para1461{"0110", 2},
			ans1461{false},
		},

		{
			para1461{"0000000001011100", 4},
			ans1461{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1461------------------------\n")

	for _, q := range qs {
		_, p := q.ans1461, q.para1461
		fmt.Printf("【input】:%v       【output】:%v\n", p, hasAllCodes(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
