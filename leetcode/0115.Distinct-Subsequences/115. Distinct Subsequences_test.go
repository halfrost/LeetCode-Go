package leetcode

import (
	"fmt"
	"testing"
)

type question115 struct {
	para115
	ans115
}

// para 是参数
// one 代表第一个参数
type para115 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans115 struct {
	one int
}

func Test_Problem115(t *testing.T) {

	qs := []question115{

		{
			para115{"rabbbit", "rabbit"},
			ans115{3},
		},

		{
			para115{"babgbag", "bag"},
			ans115{5},
		},

		{
			para115{"ab", "abc"},
			ans115{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 115------------------------\n")

	for _, q := range qs {
		a, p := q.ans115, q.para115
		fmt.Printf("【input】:%v       【output】:%v\n", p, numDistinct(p.s, p.t))
		if got := numDistinct1(p.s, p.t); got != a.one {
			t.Fatalf("numDistinct1(%q, %q) = %d, want %d", p.s, p.t, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
