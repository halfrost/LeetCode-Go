package leetcode

import (
	"fmt"
	"testing"
)

type question405 struct {
	para405
	ans405
}

// para 是参数
// one 代表第一个参数
type para405 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans405 struct {
	one string
}

func Test_Problem405(t *testing.T) {

	qs := []question405{

		{
			para405{26},
			ans405{"1a"},
		},

		{
			para405{-1},
			ans405{"ffffffff"},
		},

		{
			para405{0},
			ans405{"0"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 405------------------------\n")

	for _, q := range qs {
		a, p := q.ans405, q.para405
		fmt.Printf("【input】:%v       【output】:%v\n", p, toHex(p.one))
		if got := toHex(p.one); got != a.one {
			t.Fatalf("toHex(%d) = %q, want %q", p.one, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
