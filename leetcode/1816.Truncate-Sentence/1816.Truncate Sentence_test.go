package leetcode

import (
	"fmt"
	"testing"
)

type question1816 struct {
	para1816
	ans1816
}

// para 是参数
type para1816 struct {
	s string
	k int
}

// ans 是答案
type ans1816 struct {
	ans string
}

func Test_Problem1816(t *testing.T) {

	qs := []question1816{

		{
			para1816{"Hello how are you Contestant", 4},
			ans1816{"Hello how are you"},
		},

		{
			para1816{"What is the solution to this problem", 4},
			ans1816{"What is the solution"},
		},

		{
			para1816{"chopper is not a tanuki", 5},
			ans1816{"chopper is not a tanuki"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1816------------------------\n")

	for _, q := range qs {
		_, p := q.ans1816, q.para1816
		fmt.Printf("【input】:%v    【output】:%v\n", p, truncateSentence(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
