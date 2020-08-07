package leetcode

import (
	"fmt"
	"testing"
)

type question1175 struct {
	para1175
	ans1175
}

// para 是参数
// one 代表第一个参数
type para1175 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1175 struct {
	one int
}

func Test_Problem1175(t *testing.T) {

	qs := []question1175{

		question1175{
			para1175{5},
			ans1175{12},
		},

		question1175{
			para1175{99},
			ans1175{75763854},
		},

		question1175{
			para1175{100},
			ans1175{682289015},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1175------------------------\n")

	for _, q := range qs {
		_, p := q.ans1175, q.para1175
		fmt.Printf("【input】:%v       【output】:%v\n", p, numPrimeArrangements(p.one))
	}
	fmt.Printf("\n\n\n")
}
