package leetcode

import (
	"fmt"
	"testing"
)

type question390 struct {
	para390
	ans390
}

// para 是参数
// one 代表第一个参数
type para390 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans390 struct {
	one int
}

func Test_Problem390(t *testing.T) {

	qs := []question390{

		{
			para390{9},
			ans390{6},
		},

		{
			para390{1},
			ans390{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 390------------------------\n")

	for _, q := range qs {
		_, p := q.ans390, q.para390
		fmt.Printf("【input】:%v       【output】:%v\n", p, lastRemaining(p.n))
	}
	fmt.Printf("\n\n\n")
}
