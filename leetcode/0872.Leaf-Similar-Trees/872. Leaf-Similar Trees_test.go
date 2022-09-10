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
	}

	fmt.Printf("------------------------Leetcode Problem 872------------------------\n")

	for _, q := range qs {
		_, p := q.ans872, q.para872
		tree1 := structures.Ints2TreeNode(p.one)
		tree2 := structures.Ints2TreeNode(p.two)
		fmt.Printf("【input】:%v       【output】:%v\n", p, leafSimilar(tree1, tree2))
	}
	fmt.Printf("\n\n\n")
}
