package leetcode

import (
	"fmt"
	"testing"
)

type question169 struct {
	para169
	ans169
}

// para 是参数
// one 代表第一个参数
type para169 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans169 struct {
	one int
}

func Test_Problem169(t *testing.T) {

	qs := []question169{

		{
			para169{[]int{2, 2, 1}},
			ans169{2},
		},

		{
			para169{[]int{3, 2, 3}},
			ans169{3},
		},

		{
			para169{[]int{2, 2, 1, 1, 1, 2, 2}},
			ans169{2},
		},

		{
			para169{[]int{-2147483648}},
			ans169{-2147483648},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 169------------------------\n")

	for _, q := range qs {
		a, p := q.ans169, q.para169
		got := majorityElement(p.s)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("majorityElement(%v) = %v, want %v", p.s, got, a.one)
		}
		if got1 := majorityElement1(p.s); got1 != a.one {
			t.Fatalf("majorityElement1(%v) = %v, want %v", p.s, got1, a.one)
		}
	}

	// 当不存在多数元素时（例如空切片），解法二返回 0
	if got := majorityElement1([]int{}); got != 0 {
		t.Fatalf("majorityElement1([]) = %v, want 0", got)
	}
	fmt.Printf("\n\n\n")
}
