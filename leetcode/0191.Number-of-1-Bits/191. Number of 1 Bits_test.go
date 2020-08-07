package leetcode

import (
	"fmt"
	"testing"
)

type question191 struct {
	para191
	ans191
}

// para 是参数
// one 代表第一个参数
type para191 struct {
	one uint32
}

// ans 是答案
// one 代表第一个答案
type ans191 struct {
	one int
}

func Test_Problem191(t *testing.T) {

	qs := []question191{

		question191{
			para191{5},
			ans191{1},
		},
		question191{
			para191{13},
			ans191{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 191------------------------\n")

	for _, q := range qs {
		_, p := q.ans191, q.para191
		fmt.Printf("【input】:%v       【output】:%v\n", p, hammingWeight(p.one))
	}
	fmt.Printf("\n\n\n")
}
