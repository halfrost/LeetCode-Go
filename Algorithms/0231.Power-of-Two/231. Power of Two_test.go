package leetcode

import (
	"fmt"
	"testing"
)

type question231 struct {
	para231
	ans231
}

// para 是参数
// one 代表第一个参数
type para231 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans231 struct {
	one bool
}

func Test_Problem231(t *testing.T) {

	qs := []question231{

		question231{
			para231{1},
			ans231{true},
		},

		question231{
			para231{16},
			ans231{true},
		},

		question231{
			para231{218},
			ans231{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 231------------------------\n")

	for _, q := range qs {
		_, p := q.ans231, q.para231
		fmt.Printf("【input】:%v       【output】:%v\n", p, isPowerOfTwo(p.one))
	}
	fmt.Printf("\n\n\n")
}
