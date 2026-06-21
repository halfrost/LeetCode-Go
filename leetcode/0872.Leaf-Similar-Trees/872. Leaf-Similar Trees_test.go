package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question872 struct {
	para872
	ans872
}

// para 是参数
// one 代表第一个参数
type para872 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans872 struct {
	one bool
}

func Test_Problem872(t *testing.T) {

	qs := []question872{

		{
			para872{[]int{-10, -3, 0, 5, 9}, []int{-10, -3, 0, 5, 9}},
			ans872{true},
		},
		{
			para872{[]int{1, 2, 3}, []int{1, 3, 2}},
			ans872{false},
		},
		{
			para872{[]int{1, 2}, []int{1, 2, 3}},
			ans872{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 872------------------------\n")

	for _, q := range qs {
		a, p := q.ans872, q.para872
		tree1 := structures.Ints2TreeNode(p.one)
		tree2 := structures.Ints2TreeNode(p.two)
		got := leafSimilar(tree1, tree2)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
