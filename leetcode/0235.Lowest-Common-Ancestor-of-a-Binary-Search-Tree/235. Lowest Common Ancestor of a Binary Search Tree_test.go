package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question235 struct {
	para235
	ans235
}

// para 是参数
// one 代表第一个参数
type para235 struct {
	one []int
	two []int
	thr []int
}

// ans 是答案
// one 代表第一个答案
type ans235 struct {
	one []int
}

func Test_Problem235(t *testing.T) {

	qs := []question235{

		{
			para235{[]int{}, []int{}, []int{}},
			ans235{[]int{}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ans235{[]int{6}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans235{[]int{2}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{7}, []int{9}},
			ans235{[]int{8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 235------------------------\n")

	for _, q := range qs {
		a, p := q.ans235, q.para235
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		rootTwo := structures.Ints2TreeNode(p.two)
		rootThr := structures.Ints2TreeNode(p.thr)
		got := lowestCommonAncestor(rootOne, rootTwo, rootThr)
		fmt.Printf("【output】:%v      \n", got)
		if len(a.one) == 0 {
			if got != nil {
				t.Fatalf("input %v: expected nil, got %v", p, got)
			}
			continue
		}
		if got == nil || got.Val != a.one[0] {
			t.Fatalf("input %v: expected %v, got %v", p, a.one[0], got)
		}
	}
	fmt.Printf("\n\n\n")
}
