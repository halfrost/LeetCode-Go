package leetcode

import (
	"fmt"
	"testing"
)

type question1665 struct {
	para1665
	ans1665
}

// para 是参数
// one 代表第一个参数
type para1665 struct {
	sequence string
	word     string
}

// ans 是答案
// one 代表第一个答案
type ans1665 struct {
	one int
}

func Test_Problem1665(t *testing.T) {

	qs := []question1665{

		{
			para1665{"ababc", "ab"},
			ans1665{2},
		},

		{
			para1665{"ababc", "ba"},
			ans1665{1},
		},

		{
			para1665{"ababc", "ac"},
			ans1665{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1665------------------------\n")

	for _, q := range qs {
		_, p := q.ans1665, q.para1665
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxRepeating(p.sequence, p.word))
	}
	fmt.Printf("\n\n\n")
}
