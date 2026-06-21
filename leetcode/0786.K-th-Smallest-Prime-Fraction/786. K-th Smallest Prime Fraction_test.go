package leetcode

import (
	"fmt"
	"testing"
)

type question786 struct {
	para786
	ans786
}

// para 是参数
// one 代表第一个参数
type para786 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans786 struct {
	one []int
}

func Test_Problem786(t *testing.T) {

	qs := []question786{

		{
			para786{[]int{1, 2, 3, 5}, 3},
			ans786{[]int{2, 5}},
		},

		{
			para786{[]int{1, 7}, 1},
			ans786{[]int{1, 7}},
		},

		{
			para786{[]int{1, 2}, 1},
			ans786{[]int{1, 2}},
		},

		{
			para786{[]int{1, 2, 3, 5, 7}, 6},
			ans786{[]int{3, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 786------------------------\n")

	for _, q := range qs {
		a, p := q.ans786, q.para786
		fmt.Printf("【input】:%v       【output】:%v\n", p, kthSmallestPrimeFraction(p.A, p.K))
		got := kthSmallestPrimeFraction1(p.A, p.K)
		if got[0] != a.one[0] || got[1] != a.one[1] {
			t.Fatalf("kthSmallestPrimeFraction1(%v, %d) = %v, want %v", p.A, p.K, got, a.one)
		}
	}

	// 覆盖暴力解法的边界分支：空输入或 K 超过分数总数时返回空切片
	if got := kthSmallestPrimeFraction1([]int{}, 1); len(got) != 0 {
		t.Fatalf("kthSmallestPrimeFraction1([], 1) = %v, want []", got)
	}
	if got := kthSmallestPrimeFraction1([]int{1, 2}, 5); len(got) != 0 {
		t.Fatalf("kthSmallestPrimeFraction1([1 2], 5) = %v, want []", got)
	}

	fmt.Printf("\n\n\n")
}
