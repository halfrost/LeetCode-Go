package leetcode

import (
	"fmt"
	"testing"
)

type question1641 struct {
	para1641
	ans1641
}

// para 是参数
// one 代表第一个参数
type para1641 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1641 struct {
	one int
}

func Test_Problem1641(t *testing.T) {

	qs := []question1641{

		{
			para1641{1},
			ans1641{5},
		},

		{
			para1641{2},
			ans1641{15},
		},

		{
			para1641{33},
			ans1641{66045},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1641------------------------\n")

	for _, q := range qs {
		_, p := q.ans1641, q.para1641
		fmt.Printf("【input】:%v      【output】:%v      \n", p, countVowelStrings(p.n))
	}
	fmt.Printf("\n\n\n")
}
