package leetcode

import (
	"fmt"
	"testing"
)

type question991 struct {
	para991
	ans991
}

// para 是参数
// one 代表第一个参数
type para991 struct {
	X int
	Y int
}

// ans 是答案
// one 代表第一个答案
type ans991 struct {
	one int
}

func Test_Problem991(t *testing.T) {

	qs := []question991{

		{
			para991{2, 3},
			ans991{2},
		},

		{
			para991{5, 8},
			ans991{2},
		},

		{
			para991{3, 10},
			ans991{3},
		},

		{
			para991{1024, 1},
			ans991{1023},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 991------------------------\n")

	for _, q := range qs {
		_, p := q.ans991, q.para991
		fmt.Printf("【input】:%v       【output】:%v\n", p, brokenCalc(p.X, p.Y))
	}
	fmt.Printf("\n\n\n")
}
