package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question2096 struct {
	para2096
	ans2096
}

// para 是参数
// one 代表第一个参数
type para2096 struct {
	one        []int
	startValue int
	destValue  int
}

// ans 是答案
// one 代表第一个答案
type ans2096 struct {
	one string
}

func Test_Problem2096(t *testing.T) {

	qs := []question2096{

		{
			para2096{[]int{5, 1, 2, 3, structures.NULL, 6, 4}, 3, 6},
			ans2096{"UURL"},
		},

		{
			para2096{[]int{2, 1}, 2, 1},
			ans2096{"L"},
		},

		{
			para2096{[]int{1, 2}, 2, 1},
			ans2096{"U"},
		},

		{
			para2096{[]int{3, 1, 2}, 2, 1},
			ans2096{"UL"},
		},

		{
			para2096{[]int{7, 8, 3, 1, structures.NULL, 4, 5, 6, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 2}, 7, 5},
			ans2096{"RR"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2096------------------------\n")

	for _, q := range qs {
		_, p := q.ans2096, q.para2096
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", getDirections(root, p.startValue, p.destValue))
	}
	fmt.Printf("\n\n\n")
}
