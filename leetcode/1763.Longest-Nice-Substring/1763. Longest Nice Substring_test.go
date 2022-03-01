package leetcode

import (
	"fmt"
	"testing"
)

type question1763 struct {
	para1763
	ans1763
}

// para 是参数
// one 代表第一个参数
type para1763 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1763 struct {
	one string
}

func Test_Problem1763(t *testing.T) {

	qs := []question1763{

		{
			para1763{"YazaAay"},
			ans1763{"aAa"},
		},

		{
			para1763{"Bb"},
			ans1763{"Bb"},
		},

		{
			para1763{"c"},
			ans1763{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1763------------------------\n")

	for _, q := range qs {
		_, p := q.ans1763, q.para1763
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestNiceSubstring(p.s))
	}
	fmt.Printf("\n\n\n")
}
