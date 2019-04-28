package leetcode

import (
	"fmt"
	"testing"
)

type question263 struct {
	para263
	ans263
}

// para 是参数
// one 代表第一个参数
type para263 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans263 struct {
	one bool
}

func Test_Problem263(t *testing.T) {

	qs := []question263{

		question263{
			para263{6},
			ans263{true},
		},

		question263{
			para263{8},
			ans263{true},
		},

		question263{
			para263{14},
			ans263{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 263------------------------\n")

	for _, q := range qs {
		_, p := q.ans263, q.para263
		fmt.Printf("【input】:%v       【output】:%v\n", p, isUgly(p.one))
	}
	fmt.Printf("\n\n\n")
}
