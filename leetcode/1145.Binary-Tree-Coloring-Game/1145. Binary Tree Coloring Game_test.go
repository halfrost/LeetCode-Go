package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1145 struct {
	para1145
	ans1145
}

// para 是参数
// one 代表第一个参数
type para1145 struct {
	root []int
	n    int
	x    int
}

// ans 是答案
// one 代表第一个答案
type ans1145 struct {
	one bool
}

func Test_Problem1145(t *testing.T) {

	qs := []question1145{

		{
			para1145{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 11, 3},
			ans1145{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1145------------------------\n")

	for _, q := range qs {
		_, p := q.ans1145, q.para1145
		tree := structures.Ints2TreeNode(p.root)
		fmt.Printf("【input】:%v       【output】:%v\n", p, btreeGameWinningMove(tree, p.n, p.x))
	}
	fmt.Printf("\n\n\n")
}
