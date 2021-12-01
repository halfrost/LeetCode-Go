package leetcode

import (
	"fmt"
	"testing"
)

type question1446 struct {
	para1446
	ans1446
}

// para 是参数
type para1446 struct {
	s string
}

// ans 是答案
type ans1446 struct {
	ans int
}

func Test_Problem1446(t *testing.T) {

	qs := []question1446{

		{
			para1446{"leetcode"},
			ans1446{2},
		},

		{
			para1446{"abbcccddddeeeeedcba"},
			ans1446{5},
		},

		{
			para1446{"triplepillooooow"},
			ans1446{5},
		},

		{
			para1446{"hooraaaaaaaaaaay"},
			ans1446{11},
		},

		{
			para1446{"tourist"},
			ans1446{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1446------------------------\n")

	for _, q := range qs {
		_, p := q.ans1446, q.para1446
		fmt.Printf("【input】:%v    【output】:%v\n", p.s, maxPower(p.s))
	}
	fmt.Printf("\n\n\n")
}
