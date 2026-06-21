package leetcode

import (
	"fmt"
	"testing"
)

type question718 struct {
	para718
	ans718
}

// para 是参数
// one 代表第一个参数
type para718 struct {
	A []int
	B []int
}

// ans 是答案
// one 代表第一个答案
type ans718 struct {
	one int
}

func Test_Problem718(t *testing.T) {

	qs := []question718{

		{
			para718{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
			ans718{5},
		},

		{
			para718{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}},
			ans718{3},
		},

		{
			para718{[]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}},
			ans718{4},
		},

		{
			para718{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}},
			ans718{59},
		},

		// len(A) > len(B) covers min's return b; the pair [1,0] in A and
		// [0,16777619] in B collide on the length-2 Rabin-Karp hash but differ,
		// exercising hasSamePrefix's mismatch branch.
		{
			para718{[]int{1, 0, 99, 99}, []int{0, 16777619, 99}},
			ans718{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 718------------------------\n")

	for _, q := range qs {
		a, p := q.ans718, q.para718
		got := findLength(p.A, p.B)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("findLength(%v, %v) = %v, want %v", p.A, p.B, got, a.one)
		}
		if got1 := findLength1(p.A, p.B); got1 != a.one {
			t.Fatalf("findLength1(%v, %v) = %v, want %v", p.A, p.B, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
