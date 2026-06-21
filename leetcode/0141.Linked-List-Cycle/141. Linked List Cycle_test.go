package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question141 struct {
	para141
	ans141
}

// para 是参数
// one 代表第一个参数
type para141 struct {
	one []int
	pos int
}

// ans 是答案
// one 代表第一个答案
type ans141 struct {
	one bool
}

func Test_Problem141(t *testing.T) {

	qs := []question141{

		{
			para141{[]int{3, 2, 0, -4}, -1},
			ans141{false},
		},

		{
			para141{[]int{1, 2}, -1},
			ans141{false},
		},

		{
			para141{[]int{1}, -1},
			ans141{false},
		},

		{
			para141{[]int{3, 2, 0, -4}, 1},
			ans141{true},
		},

		{
			para141{[]int{1, 2}, 0},
			ans141{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 141------------------------\n")

	for _, q := range qs {
		a, p := q.ans141, q.para141
		var head *ListNode
		if p.pos >= 0 {
			head = structures.Ints2ListWithCycle(p.one, p.pos)
		} else {
			head = structures.Ints2List(p.one)
		}
		got := hasCycle(head)
		if got != a.one {
			t.Fatalf("input: %v pos: %v, expected: %v, got: %v", p.one, p.pos, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
