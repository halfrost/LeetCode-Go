package leetcode

import (
	"fmt"
	"testing"
)

type question331 struct {
	para331
	ans331
}

// para 是参数
// one 代表第一个参数
type para331 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans331 struct {
	one bool
}

func Test_Problem331(t *testing.T) {

	qs := []question331{

		question331{
			para331{"9,3,4,#,#,1,#,#,2,#,6,#,#"},
			ans331{true},
		},

		question331{
			para331{"1,#"},
			ans331{false},
		},

		question331{
			para331{"9,#,#,1"},
			ans331{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 331------------------------\n")

	for _, q := range qs {
		_, p := q.ans331, q.para331
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValidSerialization(p.one))
	}
	fmt.Printf("\n\n\n")
}
