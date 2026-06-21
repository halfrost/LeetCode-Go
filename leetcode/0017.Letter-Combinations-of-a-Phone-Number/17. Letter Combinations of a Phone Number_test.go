package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

func Test_Problem17(t *testing.T) {

	qs := []question17{

		{
			para17{"23"},
			ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},

		{
			para17{""},
			ans17{[]string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 17------------------------\n")

	for _, q := range qs {
		a, p := q.ans17, q.para17
		out := letterCombinations(p.s)
		fmt.Printf("【input】:%v       【output】:%v\n", p, out)
		if len(out) != len(a.one) {
			t.Fatalf("letterCombinations(%q) = %v, want %v", p.s, out, a.one)
		}
		bt := letterCombinationsBT(p.s)
		if len(bt) != len(a.one) {
			t.Fatalf("letterCombinationsBT(%q) = %v, want %v", p.s, bt, a.one)
		}
		letterCombinations_(p.s)
	}

	// Reset the shared global state so that letterCombinations_ hits the
	// len(res) == 0 branch on a fresh call.
	res = []string{}
	letterCombinations_("2")

	fmt.Printf("\n\n\n")
}
