package leetcode

import (
	"fmt"
	"testing"
)

type question223 struct {
	para223
	ans223
}

// para 是参数
// one 代表第一个参数
type para223 struct {
	A int
	B int
	C int
	D int
	E int
	F int
	G int
	H int
}

// ans 是答案
// one 代表第一个答案
type ans223 struct {
	one int
}

func Test_Problem223(t *testing.T) {

	qs := []question223{

		question223{
			para223{-3, 0, 3, 4, 0, -1, 9, 2},
			ans223{45},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 223------------------------\n")

	for _, q := range qs {
		_, p := q.ans223, q.para223
		fmt.Printf("【input】:%v       【output】:%v\n", p, computeArea(p.A, p.B, p.C, p.D, p.E, p.F, p.G, p.H))
	}
	fmt.Printf("\n\n\n")
}
