package leetcode

import (
	"fmt"
	"testing"
)

type question693 struct {
	para693
	ans693
}

// para 是参数
// one 代表第一个参数
type para693 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans693 struct {
	one bool
}

func Test_Problem693(t *testing.T) {

	qs := []question693{

		question693{
			para693{5},
			ans693{true},
		},

		question693{
			para693{7},
			ans693{false},
		},

		question693{
			para693{11},
			ans693{false},
		},

		question693{
			para693{10},
			ans693{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 693------------------------\n")

	for _, q := range qs {
		_, p := q.ans693, q.para693
		fmt.Printf("【input】:%v       【output】:%v\n", p, hasAlternatingBits(p.one))
	}
	fmt.Printf("\n\n\n")
}
