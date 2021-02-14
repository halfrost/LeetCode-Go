package leetcode

import (
	"fmt"
	"testing"
)

type question1624 struct {
	para1624
	ans1624
}

// para 是参数
// one 代表第一个参数
type para1624 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1624 struct {
	one int
}

func Test_Problem1624(t *testing.T) {

	qs := []question1624{

		{
			para1624{"aa"},
			ans1624{0},
		},

		{
			para1624{"abca"},
			ans1624{2},
		},

		{
			para1624{"cbzxy"},
			ans1624{-1},
		},

		{
			para1624{"cabbac"},
			ans1624{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1624------------------------\n")

	for _, q := range qs {
		_, p := q.ans1624, q.para1624
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxLengthBetweenEqualCharacters(p.s))
	}
	fmt.Printf("\n\n\n")
}
