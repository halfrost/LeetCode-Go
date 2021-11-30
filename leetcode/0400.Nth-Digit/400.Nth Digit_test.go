package leetcode

import (
	"fmt"
	"testing"
)

type question400 struct {
	para400
	ans400
}

// para 是参数
type para400 struct {
	n int
}

// ans 是答案
type ans400 struct {
	ans int
}

func Test_Problem400(t *testing.T) {

	qs := []question400{

		{
			para400{3},
			ans400{3},
		},

		{
			para400{11},
			ans400{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 400------------------------\n")

	for _, q := range qs {
		_, p := q.ans400, q.para400
		fmt.Printf("【input】:%v    【output】:%v\n", p.n, findNthDigit(p.n))
	}
	fmt.Printf("\n\n\n")
}
