package leetcode

import (
	"fmt"
	"testing"
)

type question54 struct {
	para54
	ans54
}

// para 是参数
// one 代表第一个参数
type para54 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans54 struct {
	one []int
}

func Test_Problem54(t *testing.T) {

	qs := []question54{

		{
			para54{[][]int{}},
			ans54{[]int{}},
		},

		{
			para54{[][]int{{}}},
			ans54{[]int{}},
		},

		{
			para54{[][]int{{3}, {2}}},
			ans54{[]int{3, 2}},
		},

		{
			para54{[][]int{{2, 3}}},
			ans54{[]int{2, 3}},
		},

		{
			para54{[][]int{{1}}},
			ans54{[]int{1}},
		},

		{
			para54{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans54{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		{
			para54{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
			ans54{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		},
		{
			para54{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}},
			ans54{[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 54------------------------\n")

	for _, q := range qs {
		a, p := q.ans54, q.para54
		got := spiralOrder(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		got2 := spiralOrder2(p.one)
		if len(got) != len(a.one) {
			t.Fatalf("spiralOrder(%v) = %v, want %v", p.one, got, a.one)
		}
		for i := range a.one {
			if got[i] != a.one[i] {
				t.Fatalf("spiralOrder(%v) = %v, want %v", p.one, got, a.one)
			}
		}
		if len(got2) != len(a.one) {
			t.Fatalf("spiralOrder2(%v) = %v, want %v", p.one, got2, a.one)
		}
		for i := range a.one {
			if got2[i] != a.one[i] {
				t.Fatalf("spiralOrder2(%v) = %v, want %v", p.one, got2, a.one)
			}
		}
	}
	fmt.Printf("\n\n\n")
}
