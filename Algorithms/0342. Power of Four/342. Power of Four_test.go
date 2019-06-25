package leetcode

import (
	"fmt"
	"testing"
)

type question342 struct {
	para342
	ans342
}

// para 是参数
// one 代表第一个参数
type para342 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans342 struct {
	one bool
}

func Test_Problem342(t *testing.T) {

	qs := []question342{

		question342{
			para342{16},
			ans342{true},
		},

		question342{
			para342{5},
			ans342{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 342------------------------\n")

	for _, q := range qs {
		_, p := q.ans342, q.para342
		fmt.Printf("【input】:%v       【output】:%v\n", p, isPowerOfFour(p.one))
	}
	fmt.Printf("\n\n\n")
}
