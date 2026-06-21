package leetcode

import (
	"fmt"
	"testing"
)

type question198 struct {
	para198
	ans198
}

// para 是参数
// one 代表第一个参数
type para198 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans198 struct {
	one int
}

func Test_Problem198(t *testing.T) {

	qs := []question198{

		{
			para198{[]int{}},
			ans198{0},
		},

		{
			para198{[]int{5}},
			ans198{5},
		},

		{
			para198{[]int{1, 2}},
			ans198{2},
		},

		{
			para198{[]int{1, 2, 3, 1}},
			ans198{4},
		},
		{
			para198{[]int{2, 7, 9, 3, 1}},
			ans198{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 198------------------------\n")

	for _, q := range qs {
		a, p := q.ans198, q.para198
		fmt.Printf("【input】:%v       【output】:%v\n", p, rob198(p.one))
		if got := rob198(p.one); got != a.one {
			t.Fatalf("rob198(%v) = %v, want %v", p.one, got, a.one)
		}
		if got := rob(p.one); got != a.one {
			t.Fatalf("rob(%v) = %v, want %v", p.one, got, a.one)
		}
		if got := rob198_1(p.one); got != a.one {
			t.Fatalf("rob198_1(%v) = %v, want %v", p.one, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
