package leetcode

import (
	"fmt"
	"testing"
)

type question925 struct {
	para925
	ans925
}

// para 是参数
// one 代表第一个参数
type para925 struct {
	name  string
	typed string
}

// ans 是答案
// one 代表第一个答案
type ans925 struct {
	one bool
}

func Test_Problem925(t *testing.T) {

	qs := []question925{

		question925{
			para925{"alex", "aaleex"},
			ans925{true},
		},

		question925{
			para925{"saeed", "ssaaedd"},
			ans925{false},
		},

		question925{
			para925{"leelee", "lleeelee"},
			ans925{true},
		},

		question925{
			para925{"laiden", "laiden"},
			ans925{true},
		},

		question925{
			para925{"kikcxmvzi", "kiikcxxmmvvzz"},
			ans925{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 925------------------------\n")

	for _, q := range qs {
		_, p := q.ans925, q.para925
		fmt.Printf("【input】:%v       【output】:%v\n", p, isLongPressedName(p.name, p.typed))
	}
	fmt.Printf("\n\n\n")
}
