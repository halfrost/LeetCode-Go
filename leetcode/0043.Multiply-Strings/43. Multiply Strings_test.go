package leetcode

import (
	"fmt"
	"testing"
)

type question43 struct {
	para43
	ans43
}

// para 是参数
// one 代表第一个参数
type para43 struct {
	num1 string
	num2 string
}

// ans 是答案
// one 代表第一个答案
type ans43 struct {
	one string
}

func Test_Problem43(t *testing.T) {

	qs := []question43{

		{
			para43{"2", "3"},
			ans43{"6"},
		},

		{
			para43{"123", "456"},
			ans43{"56088"},
		},

		{
			para43{"0", "52"},
			ans43{"0"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 43------------------------\n")

	for _, q := range qs {
		a, p := q.ans43, q.para43
		got := multiply(p.num1, p.num2)
		if got != a.one {
			t.Fatalf("multiply(%q, %q) = %q, want %q", p.num1, p.num2, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
