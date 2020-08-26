package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question337 struct {
	para337
	ans337
}

// para 是参数
// one 代表第一个参数
type para337 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans337 struct {
	one int
}

func Test_Problem337(t *testing.T) {

	qs := []question337{

		{
			para337{[]int{3, 2, 3, structures.NULL, 3, structures.NULL, 1}},
			ans337{7},
		},

		{
			para337{[]int{}},
			ans337{0},
		},

		{
			para337{[]int{3, 4, 5, 1, 3, structures.NULL, 1}},
			ans337{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 337------------------------\n")

	for _, q := range qs {
		_, p := q.ans337, q.para337
		fmt.Printf("【input】:%v       【output】:%v\n", p, rob337(structures.Ints2TreeNode(p.one)))
	}
	fmt.Printf("\n\n\n")
}
