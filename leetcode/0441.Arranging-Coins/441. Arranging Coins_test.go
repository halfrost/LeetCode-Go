package leetcode

import (
	"fmt"
	"testing"
)

type question441 struct {
	para441
	ans441
}

// para 是参数
// one 代表第一个参数
type para441 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans441 struct {
	one int
}

func Test_Problem441(t *testing.T) {

	qs := []question441{

		{
			para441{5},
			ans441{2},
		},

		{
			para441{8},
			ans441{3},
		},

		{
			para441{0},
			ans441{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 441------------------------\n")

	for _, q := range qs {
		a, p := q.ans441, q.para441
		fmt.Printf("【input】:%v       【output】:%v\n", p, arrangeCoins(p.n))
		if got := arrangeCoins(p.n); got != a.one {
			t.Fatalf("arrangeCoins(%d) = %d, want %d", p.n, got, a.one)
		}
		if got := arrangeCoins1(p.n); got != a.one {
			t.Fatalf("arrangeCoins1(%d) = %d, want %d", p.n, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
