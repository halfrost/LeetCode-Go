package leetcode

import (
	"fmt"
	"testing"
)

type question816 struct {
	para816
	ans816
}

// para 是参数
// one 代表第一个参数
type para816 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans816 struct {
	one []string
}

func Test_Problem816(t *testing.T) {

	qs := []question816{

		{
			para816{"(120123)"},
			ans816{[]string{"(1, 20123)", " (1, 2.0123)", " (1, 20.123)", " (1, 201.23)", " (1, 2012.3)", " (12, 0.123)", " (1.2, 0.123)", " (120, 123)", " (120, 1.23)", " (120, 12.3)", " (1201, 23) ", "(1201, 2.3)", " (1.201, 23)", " (1.201, 2.3) ", "(12.01, 23)", " (12.01, 2.3) ", "(120.1, 23)", " (120.1, 2.3) ", "(12012, 3)", " (1.2012, 3)", " (12.012, 3)", " (120.12, 3)", " (1201.2, 3)"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 816------------------------\n")

	for _, q := range qs {
		_, p := q.ans816, q.para816
		fmt.Printf("【input】:%v       【output】:%v\n", p, ambiguousCoordinates(p.one))
	}
	fmt.Printf("\n\n\n")
}
