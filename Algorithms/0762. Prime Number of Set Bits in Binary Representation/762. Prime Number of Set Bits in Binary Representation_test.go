package leetcode

import (
	"fmt"
	"testing"
)

type question762 struct {
	para762
	ans762
}

// para 是参数
// one 代表第一个参数
type para762 struct {
	l int
	r int
}

// ans 是答案
// one 代表第一个答案
type ans762 struct {
	one int
}

func Test_Problem762(t *testing.T) {

	qs := []question762{

		question762{
			para762{6, 10},
			ans762{4},
		},

		question762{
			para762{10, 15},
			ans762{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 762------------------------\n")

	for _, q := range qs {
		_, p := q.ans762, q.para762
		fmt.Printf("【input】:%v       【output】:%v\n", p, countPrimeSetBits(p.l, p.r))
	}
	fmt.Printf("\n\n\n")
}
