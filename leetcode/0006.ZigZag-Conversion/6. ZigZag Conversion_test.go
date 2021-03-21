package leetcode

import (
	"fmt"
	"testing"
)

type question6 struct {
	para6
	ans6
}

// para 是参数
// one 代表第一个参数
type para6 struct {
	s       string
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans6 struct {
	one string
}

func Test_Problem6(t *testing.T) {

	qs := []question6{

		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},

		{
			para6{"PAYPALISHIRING", 4},
			ans6{"PINALSIGYAHRPI"},
		},

		{
			para6{"A", 1},
			ans6{"A"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 6------------------------\n")

	for _, q := range qs {
		_, p := q.ans6, q.para6
		fmt.Printf("【input】:%v       【output】:%v\n", p, convert(p.s, p.numRows))
	}
	fmt.Printf("\n\n\n")
}
