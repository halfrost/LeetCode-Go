package leetcode

import (
	"fmt"
	"testing"
)

type question96 struct {
	para96
	ans96
}

// para 是参数
// one 代表第一个参数
type para96 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans96 struct {
	one int
}

func Test_Problem96(t *testing.T) {

	qs := []question96{

		question96{
			para96{1},
			ans96{1},
		},

		question96{
			para96{3},
			ans96{5},
		},

		question96{
			para96{4},
			ans96{14},
		},

		question96{
			para96{5},
			ans96{42},
		},

		question96{
			para96{6},
			ans96{132},
		},

		question96{
			para96{7},
			ans96{429},
		},

		question96{
			para96{8},
			ans96{1430},
		},

		question96{
			para96{9},
			ans96{4862},
		},
		question96{
			para96{10},
			ans96{16796},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 96------------------------\n")

	for _, q := range qs {
		_, p := q.ans96, q.para96
		fmt.Printf("【input】:%v       【output】:%v\n", p, numTrees(p.one))
	}
	fmt.Printf("\n\n\n")
}
