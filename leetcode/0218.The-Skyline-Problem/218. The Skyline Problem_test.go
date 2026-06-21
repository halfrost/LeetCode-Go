package leetcode

import (
	"fmt"
	"testing"
)

type question218 struct {
	para218
	ans218
}

// para 是参数
// one 代表第一个参数
type para218 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans218 struct {
	one [][]int
}

func Test_Problem218(t *testing.T) {

	qs := []question218{

		{
			para218{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}},
			ans218{[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		},

		{
			para218{[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3}}},
			ans218{[][]int{{1, 3}, {3, 0}}},
		},

		{
			para218{[][]int{{4, 9, 10}, {4, 9, 15}, {4, 9, 12}, {10, 12, 10}, {10, 12, 8}}},
			ans218{[][]int{{4, 15}, {9, 0}, {10, 10}, {12, 0}}},
		},

		{
			para218{[][]int{}},
			ans218{[][]int{}},
		},

		{
			para218{[][]int{{1, 10, 5}, {2, 10, 6}, {3, 10, 7}, {4, 10, 8}, {5, 10, 9}, {6, 10, 4}, {7, 10, 3}}},
			ans218{[][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}, {10, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 218------------------------\n")

	for _, q := range qs {
		_, p := q.ans218, q.para218
		fmt.Printf("【input】:%v       【output】:%v\n", p, getSkyline(p.one))
		getSkyline1(p.one)
		getSkyline2(p.one)
	}
	fmt.Printf("\n\n\n")
}

func Test_IndexMaxPQ218(t *testing.T) {
	// Build a heap large enough that removing the top forces the moved
	// element to sink through multiple levels, exercising both the
	// right-child selection (k++) branch and the exch/n=k body in sink.
	pq := NewIndexMaxPQ(8)
	vals := []int{1, 8, 7, 6, 5, 4, 3, 2}
	for i, v := range vals {
		pq.Enque(i, v)
	}
	// Remove the current maximum repeatedly; each removal moves a small
	// element to the root and sinks it down.
	pq.Remove(1) // remove key with value 8 (the max)
	if got := pq.Front(); got != 7 {
		t.Fatalf("after removing max(8), Front()=%d, want 7", got)
	}
	pq.Remove(2) // remove value 7
	if got := pq.Front(); got != 6 {
		t.Fatalf("after removing 7, Front()=%d, want 6", got)
	}
	fmt.Printf("IndexMaxPQ sink covered, Front=%d\n", pq.Front())
}
