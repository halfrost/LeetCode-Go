package leetcode

import (
	"fmt"
	"testing"
)

type question319 struct {
	para319
	ans319
}

// para 是参数
type para319 struct {
	n int
}

// ans 是答案
type ans319 struct {
	ans int
}

func Test_Problem319(t *testing.T) {

	qs := []question319{

		{
			para319{3},
			ans319{1},
		},

		{
			para319{0},
			ans319{0},
		},

		{
			para319{1},
			ans319{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 319------------------------\n")

	for _, q := range qs {
		_, p := q.ans319, q.para319
		fmt.Printf("【input】:%v       【output】:%v\n", p, bulbSwitch(p.n))
	}
	fmt.Printf("\n\n\n")
}
