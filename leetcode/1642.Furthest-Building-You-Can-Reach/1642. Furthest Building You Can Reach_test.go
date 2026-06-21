package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

type question1642 struct {
	para1642
	ans1642
}

// para 是参数
// one 代表第一个参数
type para1642 struct {
	heights []int
	bricks  int
	ladders int
}

// ans 是答案
// one 代表第一个答案
type ans1642 struct {
	one int
}

func Test_Problem1642(t *testing.T) {

	qs := []question1642{

		{
			para1642{[]int{1, 5, 1, 2, 3, 4, 10000}, 4, 1},
			ans1642{5},
		},

		{
			para1642{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1},
			ans1642{4},
		},

		{
			para1642{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2},
			ans1642{7},
		},

		{
			para1642{[]int{14, 3, 19, 3}, 17, 0},
			ans1642{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1642------------------------\n")

	for _, q := range qs {
		a, p := q.ans1642, q.para1642
		out := furthestBuilding(p.heights, p.bricks, p.ladders)
		fmt.Printf("【input】:%v      【output】:%v      \n", p, out)
		if out != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, out)
		}
	}
	fmt.Printf("\n\n\n")

	// cover the heap Pop method, which the solution itself never invokes
	pq := &heightDiffPQ{}
	heap.Push(pq, 3)
	heap.Push(pq, 1)
	heap.Push(pq, 2)
	if got := heap.Pop(pq).(int); got != 1 {
		t.Fatalf("heap.Pop expected: 1, got: %v", got)
	}
}
