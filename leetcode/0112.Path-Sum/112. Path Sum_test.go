package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question112 struct {
	para112
	ans112
}

// para 是参数
// one 代表第一个参数
type para112 struct {
	one []int
	sum int
}

// ans 是答案
// one 代表第一个答案
type ans112 struct {
	one bool
}

func Test_Problem112(t *testing.T) {

	qs := []question112{

		{
			para112{[]int{}, 0},
			ans112{false},
		},

		{
			para112{[]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, structures.NULL, 1}, 22},
			ans112{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 112------------------------\n")

	for _, q := range qs {
		_, p := q.ans112, q.para112
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", hasPathSum(root, p.sum))
	}
	fmt.Printf("\n\n\n")
}
