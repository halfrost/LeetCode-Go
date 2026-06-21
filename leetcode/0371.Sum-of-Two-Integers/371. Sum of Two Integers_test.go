package leetcode

import (
	"fmt"
	"testing"
)

type question371 struct {
	para371
	ans371
}

// para 是参数
// one 代表第一个参数
type para371 struct {
	a int
	b int
}

// ans 是答案
// one 代表第一个答案
type ans371 struct {
	one int
}

func Test_Problem371(t *testing.T) {

	qs := []question371{

		{
			para371{1, 2},
			ans371{3},
		},

		{
			para371{-2, 3},
			ans371{1},
		},

		{
			para371{0, 5},
			ans371{5},
		},

		{
			para371{7, 0},
			ans371{7},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 371------------------------\n")

	for _, q := range qs {
		a, p := q.ans371, q.para371
		got := getSum(p.a, p.b)
		if got != a.one {
			t.Fatalf("input %v: expected %v, got %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
