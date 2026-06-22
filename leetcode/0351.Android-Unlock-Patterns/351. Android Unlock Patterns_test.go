package leetcode

import (
	"fmt"
	"testing"
)

type question351 struct {
	para351
	ans351
}

// para is the parameter
// one represents the first parameter, two the second
type para351 struct {
	one int
	two int
}

// ans is the answer
// one represents the first answer
type ans351 struct {
	one int
}

func Test_Problem351(t *testing.T) {

	qs := []question351{

		{
			para351{1, 1},
			ans351{9},
		},

		{
			para351{1, 2},
			ans351{65},
		},

		{
			para351{2, 3},
			ans351{376},
		},

		{
			para351{4, 5},
			ans351{8776},
		},

		{
			para351{4, 9},
			ans351{389112},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 351------------------------\n")

	for _, q := range qs {
		a, p := q.ans351, q.para351
		got := numberOfPatterns(p.one, p.two)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input m=%v n=%v: expected %v, got %v", p.one, p.two, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
