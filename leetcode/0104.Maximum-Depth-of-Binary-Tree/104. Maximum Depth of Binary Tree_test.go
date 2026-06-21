package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question104 struct {
	para104
	ans104
}

// para 是参数
// one 代表第一个参数
type para104 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans104 struct {
	one int
}

func Test_Problem104(t *testing.T) {

	qs := []question104{

		{
			para104{[]int{}},
			ans104{0},
		},

		{
			para104{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans104{3},
		},

		{
			para104{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, structures.NULL, 5}},
			ans104{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 104------------------------\n")

	for _, q := range qs {
		a, p := q.ans104, q.para104
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		got := maxDepth(root)
		fmt.Printf("【output】:%v      \n", got)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p.one, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
