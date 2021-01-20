package leetcode

import (
	"fmt"
	"testing"
)

type question1668 struct {
	para1668
	ans1668
}

// para 是参数
// one 代表第一个参数
type para1668 struct {
	sequence string
	word     string
}

// ans 是答案
// one 代表第一个答案
type ans1668 struct {
	one int
}

func Test_Problem1668(t *testing.T) {

	qs := []question1668{

		{
			para1668{"ababc", "ab"},
			ans1668{2},
		},

		{
			para1668{"ababc", "ba"},
			ans1668{1},
		},

		{
			para1668{"ababc", "ac"},
			ans1668{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1668------------------------\n")

	for _, q := range qs {
		_, p := q.ans1668, q.para1668
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxRepeating(p.sequence, p.word))
	}
	fmt.Printf("\n\n\n")
}
