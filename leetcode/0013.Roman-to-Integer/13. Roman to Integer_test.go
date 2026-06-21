package leetcode

import (
	"fmt"
	"testing"
)

type question13 struct {
	para13
	ans13
}

// para 是参数
// one 代表第一个参数
type para13 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans13 struct {
	one int
}

func Test_Problem13(t *testing.T) {

	qs := []question13{

		{
			para13{"III"},
			ans13{3},
		},

		{
			para13{"IV"},
			ans13{4},
		},

		{
			para13{"IX"},
			ans13{9},
		},

		{
			para13{"LVIII"},
			ans13{58},
		},

		{
			para13{"MCMXCIV"},
			ans13{1994},
		},

		{
			para13{"MCMXICIVI"},
			ans13{2014},
		},

		{
			para13{""},
			ans13{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 13------------------------\n")

	for _, q := range qs {
		a, p := q.ans13, q.para13
		got := romanToInt(p.one)
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, got)
		if got != a.one {
			t.Fatalf("input %q: got %d, want %d", p.one, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
