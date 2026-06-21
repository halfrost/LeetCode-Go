package leetcode

import (
	"fmt"
	"math"
	"testing"
)

type question29 struct {
	para29
	ans29
}

// para 是参数
// one 代表第一个参数
type para29 struct {
	dividend int
	divisor  int
}

// ans 是答案
// one 代表第一个答案
type ans29 struct {
	one int
}

func Test_Problem29(t *testing.T) {

	qs := []question29{

		{
			para29{10, 3},
			ans29{3},
		},

		{
			para29{7, -3},
			ans29{-2},
		},

		{
			para29{-1, 1},
			ans29{-1},
		},

		{
			para29{1, -1},
			ans29{-1},
		},

		{
			para29{2147483647, 3},
			ans29{715827882},
		},

		{
			para29{0, 5},
			ans29{0},
		},

		{
			para29{10, 1},
			ans29{10},
		},

		{
			para29{math.MinInt32, -1},
			ans29{math.MaxInt32},
		},

		{
			para29{math.MaxInt32 + 10, 2},
			ans29{1073741823},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 29------------------------\n")

	for _, q := range qs {
		a, p := q.ans29, q.para29
		got := divide(p.dividend, p.divisor)
		fmt.Printf("【input】:%v    【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("divide(%d, %d) = %d, want %d", p.dividend, p.divisor, got, a.one)
		}
		got1 := divide1(p.dividend, p.divisor)
		// divide caps dividend to math.MaxInt32 while divide1 does not, so only
		// cross-check the two solutions for inputs within the int32 range.
		if p.dividend <= math.MaxInt32 && got1 != a.one {
			t.Fatalf("divide1(%d, %d) = %d, want %d", p.dividend, p.divisor, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
