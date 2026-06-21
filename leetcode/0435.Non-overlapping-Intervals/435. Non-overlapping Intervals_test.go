package leetcode

import (
	"fmt"
	"testing"
)

type question435 struct {
	para435
	ans435
}

// para 是参数
// one 代表第一个参数
type para435 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans435 struct {
	one int
}

func Test_Problem435(t *testing.T) {

	qs := []question435{

		{
			para435{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			ans435{1},
		},

		{
			para435{[][]int{{1, 2}, {1, 2}, {1, 2}}},
			ans435{2},
		},

		{
			para435{[][]int{{1, 2}, {2, 3}}},
			ans435{0},
		},

		{
			para435{[][]int{}},
			ans435{0},
		},

		{
			para435{[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}},
			ans435{2},
		},

		{
			para435{[][]int{{1, 10}, {2, 3}}},
			ans435{1},
		},

		{
			para435{[][]int{{1, 2}, {3, 4}, {5, 6}, {1, 6}}},
			ans435{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 435------------------------\n")

	for _, q := range qs {
		a, p := q.ans435, q.para435
		got1 := eraseOverlapIntervals1(clone435(p.one))
		got2 := eraseOverlapIntervals(clone435(p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, got1)
		if got1 != a.one {
			t.Fatalf("eraseOverlapIntervals1(%v) = %v, want %v", p.one, got1, a.one)
		}
		if got2 != a.one {
			t.Fatalf("eraseOverlapIntervals(%v) = %v, want %v", p.one, got2, a.one)
		}
	}
	if got := max(2, 1); got != 2 {
		t.Fatalf("max(2, 1) = %v, want 2", got)
	}
	if got := max(1, 2); got != 2 {
		t.Fatalf("max(1, 2) = %v, want 2", got)
	}
	fmt.Printf("\n\n\n")
}

func clone435(intervals [][]int) [][]int {
	res := make([][]int, len(intervals))
	for i, v := range intervals {
		res[i] = append([]int{}, v...)
	}
	return res
}
