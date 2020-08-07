package leetcode

import (
	"fmt"
	"testing"
)

type question470 struct {
	para470
	ans470
}

// para 是参数
// one 代表第一个参数
type para470 struct {
}

// ans 是答案
// one 代表第一个答案
type ans470 struct {
	one int
}

func Test_Problem470(t *testing.T) {

	qs := []question470{

		question470{
			para470{},
			ans470{2},
		},

		question470{
			para470{},
			ans470{0},
		},

		question470{
			para470{},
			ans470{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 470------------------------\n")

	for _, q := range qs {
		_, p := q.ans470, q.para470
		fmt.Printf("【input】:%v       【output】:%v\n", p, rand10())
	}
	fmt.Printf("\n\n\n")
}
