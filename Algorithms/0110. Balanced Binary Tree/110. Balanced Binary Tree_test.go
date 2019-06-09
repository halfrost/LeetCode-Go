package leetcode

import (
	"fmt"
	"testing"
)

type question110 struct {
	para110
	ans110
}

// para 是参数
// one 代表第一个参数
type para110 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans110 struct {
	one bool
}

func Test_Problem110(t *testing.T) {

	qs := []question110{

		question110{
			para110{[]int{3, 4, 4, 5, NULL, NULL, 5, 6, NULL, NULL, 6}},
			ans110{false},
		},

		question110{
			para110{[]int{1, 2, 2, NULL, 3, 3}},
			ans110{true},
		},

		question110{
			para110{[]int{}},
			ans110{true},
		},

		question110{
			para110{[]int{1}},
			ans110{true},
		},

		question110{
			para110{[]int{1, 2, 3}},
			ans110{true},
		},

		question110{
			para110{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans110{true},
		},

		question110{
			para110{[]int{1, 2, 2, NULL, 3, NULL, 3}},
			ans110{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 110------------------------\n")

	for _, q := range qs {
		_, p := q.ans110, q.para110
		fmt.Printf("【input】:%v      ", p)
		rootOne := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isBalanced(rootOne))
	}
	fmt.Printf("\n\n\n")
}
