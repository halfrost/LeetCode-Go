package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question142 struct {
	para142
	ans142
}

type para142 struct {
	nums []int
	pos  int
}

type ans142 struct {
	hasCycle bool
	val      int
}

func Test_Problem142(t *testing.T) {
	qs := []question142{
		{
			para142{[]int{3, 2, 0, -4}, 1},
			ans142{true, 2},
		},
		{
			para142{[]int{1, 2}, 0},
			ans142{true, 1},
		},
		{
			para142{[]int{1}, -1},
			ans142{false, 0},
		},
		{
			para142{[]int{1, 2, 3}, -1},
			ans142{false, 0},
		},
		{
			para142{[]int{}, -1},
			ans142{false, 0},
		},
	}

	for _, q := range qs {
		a, p := q.ans142, q.para142
		head := structures.Ints2ListWithCycle(p.nums, p.pos)
		node := detectCycle(head)
		fmt.Printf("nums = %v pos = %v node = %v\n", p.nums, p.pos, node)
		if a.hasCycle {
			if node == nil {
				t.Fatalf("expected cycle entry node with val %d, got nil", a.val)
			}
			if node.Val != a.val {
				t.Fatalf("expected cycle entry val %d, got %d", a.val, node.Val)
			}
		} else {
			if node != nil {
				t.Fatalf("expected no cycle (nil), got node with val %d", node.Val)
			}
		}
	}
}
