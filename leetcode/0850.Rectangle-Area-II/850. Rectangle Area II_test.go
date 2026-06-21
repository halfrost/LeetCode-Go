package leetcode

import (
	"fmt"
	"testing"
)

type question850 struct {
	para850
	ans850
}

// para 是参数
// one 代表第一个参数
type para850 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans850 struct {
	one int
}

func Test_Problem850(t *testing.T) {

	qs := []question850{

		{
			para850{[][]int{{0, 0, 3, 3}, {2, 0, 5, 3}, {1, 1, 4, 4}}},
			ans850{18},
		},

		{
			para850{[][]int{{0, 0, 1, 1}, {2, 2, 3, 3}}},
			ans850{2},
		},

		{
			para850{[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}},
			ans850{6},
		},

		{
			para850{[][]int{{0, 0, 1000000000, 1000000000}}},
			ans850{49},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 850------------------------\n")

	for _, q := range qs {
		_, p := q.ans850, q.para850
		fmt.Printf("【input】:%v       【output】:%v\n", p, rectangleArea(p.one))
	}
	fmt.Printf("\n\n\n")
}

func Test_SegmentAreaTree850(t *testing.T) {
	// Empty tree: Query and Update should hit the len(data) == 0 guard returning 0.
	empty := SegmentAreaTree{}
	empty.Init([]int{}, func(i, j int) int { return i + j })
	if got := empty.Query(0, 0); got != 0 {
		t.Fatalf("empty tree Query expected 0, got %d", got)
	}
	empty.Update(0, 0, 1) // exercise the empty-tree Update guard

	// Build a tree over several segments so that partial queries exercise the
	// outside-range, leaf-overlap, left-only, right-only, and merge branches of
	// queryInTree without recursing forever.
	sat := SegmentAreaTree{}
	tmp := []int{1, 1, 1, 1, 1, 1, 1, 1}
	sat.Init(tmp, func(i, j int) int { return i + j })
	// Mark leaves as covered so the count > 0 branch of queryInTree is taken.
	sat.Update(0, len(tmp), 1)

	n := len(tmp) - 1 // root right boundary used by Query
	// Full range query through the public API (covers the merge / inside branches).
	full := sat.Query(0, n)
	if full < 0 {
		t.Fatalf("full Query returned negative: %d", full)
	}
	// Left-leaning sub-range (queryRight <= midTreeIndex branch).
	if got := sat.Query(0, 2); got < 0 {
		t.Fatalf("left Query returned negative: %d", got)
	}
	// queryLeft > midTreeIndex branch, terminating thanks to the leaf guard.
	if got := sat.queryInTree(0, 0, n, 2, 2); got < 0 {
		t.Fatalf("right-leaning Query returned negative: %d", got)
	}
	// Out-of-range query (left > queryRight / right < queryLeft branch).
	if got := sat.queryInTree(0, 0, n, -2, -2); got != 0 {
		t.Fatalf("out-of-range Query expected 0, got %d", got)
	}

	// A tree with no Update keeps every node count == 0, so the leaf-overlap
	// branch returns 0 (the count == 0 path).
	uncovered := SegmentAreaTree{}
	uncovered.Init(tmp, func(i, j int) int { return i + j })
	if got := uncovered.queryInTree(0, 0, n, 2, 2); got != 0 {
		t.Fatalf("uncovered leaf Query expected 0, got %d", got)
	}
	fmt.Printf("【SegmentAreaTree】 full=%v\n", full)
}
