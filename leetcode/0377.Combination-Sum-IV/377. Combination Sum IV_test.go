package leetcode

import (
	"fmt"
	"testing"
)

type question377 struct {
	para377
	ans377
}

// para 是参数
// one 代表第一个参数
type para377 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans377 struct {
	one int
}

func Test_Problem377(t *testing.T) {

	qs := []question377{

		{
			para377{[]int{1, 2, 3}, 4},
			ans377{7},
		},

		{
			para377{[]int{1, 2, 3}, 32},
			ans377{181997601},
		},

		{
			para377{[]int{}, 4},
			ans377{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 377------------------------\n")

	for _, q := range qs {
		a, p := q.ans377, q.para377
		fmt.Printf("【input】:%v       【output】:%v\n", p, combinationSum4(p.n, p.k))
		if got := combinationSum4(p.n, p.k); got != a.one {
			t.Fatalf("combinationSum4(%v, %v) = %v, want %v", p.n, p.k, got, a.one)
		}
		// 暴力解法在 target 较大时超时，仅对小输入调用以覆盖代码。
		if p.k <= 4 {
			if got := combinationSum41(p.n, p.k); got != a.one {
				t.Fatalf("combinationSum41(%v, %v) = %v, want %v", p.n, p.k, got, a.one)
			}
		}
	}
	fmt.Printf("\n\n\n")
}
