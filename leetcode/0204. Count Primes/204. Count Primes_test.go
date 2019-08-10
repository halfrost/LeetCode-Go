package leetcode

import (
	"fmt"
	"testing"
)

type question204 struct {
	para204
	ans204
}

// para 是参数
// one 代表第一个参数
type para204 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans204 struct {
	one int
}

func Test_Problem204(t *testing.T) {

	qs := []question204{

		question204{
			para204{10},
			ans204{4},
		},

		question204{
			para204{100},
			ans204{25},
		},

		question204{
			para204{1000},
			ans204{168},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 204------------------------\n")

	for _, q := range qs {
		_, p := q.ans204, q.para204
		fmt.Printf("【input】:%v       【output】:%v\n", p, countPrimes(p.one))
	}
	fmt.Printf("\n\n\n")
}
