package leetcode

import (
	"fmt"
	"testing"
)

type question828 struct {
	para828
	ans828
}

// para 是参数
// one 代表第一个参数
type para828 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans828 struct {
	one int
}

func Test_Problem828(t *testing.T) {

	qs := []question828{

		question828{
			para828{"BABABBABAA"},
			ans828{35},
		},

		question828{
			para828{"ABC"},
			ans828{10},
		},

		question828{
			para828{"ABA"},
			ans828{8},
		},

		question828{
			para828{"ABAB"},
			ans828{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 828------------------------\n")

	for _, q := range qs {
		_, p := q.ans828, q.para828
		fmt.Printf("【input】:%v       【output】:%v\n", p, uniqueLetterString(p.one))
	}
	fmt.Printf("\n\n\n")
}
