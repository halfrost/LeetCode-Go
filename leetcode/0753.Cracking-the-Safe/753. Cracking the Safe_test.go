package leetcode

import (
	"fmt"
	"testing"
)

type question753 struct {
	para753
	ans753
}

// para 是参数
// one 代表第一个参数
type para753 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans753 struct {
	one string
}

func Test_Problem753(t *testing.T) {

	qs := []question753{

		{
			para753{1, 2},
			ans753{"01"},
		},

		{
			para753{2, 2},
			ans753{"00110"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 753------------------------\n")

	for _, q := range qs {
		_, p := q.ans753, q.para753
		fmt.Printf("【input】:%v       【output】:%v\n", p, crackSafe(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
