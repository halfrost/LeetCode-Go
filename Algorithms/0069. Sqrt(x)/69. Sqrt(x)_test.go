package leetcode

import (
	"fmt"
	"testing"
)

type question69 struct {
	para69
	ans69
}

// para 是参数
// one 代表第一个参数
type para69 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans69 struct {
	one int
}

func Test_Problem69(t *testing.T) {

	qs := []question69{

		question69{
			para69{4},
			ans69{2},
		},

		question69{
			para69{8},
			ans69{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 69------------------------\n")

	for _, q := range qs {
		_, p := q.ans69, q.para69
		fmt.Printf("【input】:%v       【output】:%v\n", p, mySqrt1(p.one))
	}
	fmt.Printf("\n\n\n")
}
