package leetcode

import (
	"fmt"
	"testing"
)

type question150 struct {
	para150
	ans150
}

// para 是参数
// one 代表第一个参数
type para150 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans150 struct {
	one int
}

func Test_Problem150(t *testing.T) {

	qs := []question150{

		question150{
			para150{[]string{"18"}},
			ans150{18},
		},

		question150{
			para150{[]string{"2", "1", "+", "3", "*"}},
			ans150{9},
		},
		question150{
			para150{[]string{"4", "13", "5", "/", "+"}},
			ans150{6},
		},

		question150{
			para150{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			ans150{22},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 150------------------------\n")

	for _, q := range qs {
		_, p := q.ans150, q.para150
		fmt.Printf("【input】:%v       【output】:%v\n", p, evalRPN(p.one))
	}
	fmt.Printf("\n\n\n")
}
