package leetcode

import (
	"fmt"
	"testing"
)

type question897 struct {
	para897
	ans897
}

// para 是参数
// one 代表第一个参数
type para897 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans897 struct {
	one []int
}

func Test_Problem897(t *testing.T) {

	qs := []question897{

		question897{
			para897{[]int{5, 3, 6, 2, 4, NULL, 8, 1, NULL, NULL, NULL, 7, 9}},
			ans897{[]int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5, NULL, 6, NULL, 7, NULL, 8, NULL, 9}},
		},

		question897{
			para897{[]int{3, 4, 4, 5, NULL, NULL, 5, 6, NULL, NULL, 6}},
			ans897{[]int{6, NULL, 5, NULL, 4, NULL, 3, NULL, 4, NULL, 5, NULL, 6}},
		},

		question897{
			para897{[]int{1, 2, 2, NULL, 3, 3}},
			ans897{[]int{2, NULL, 3, NULL, 1, NULL, 3, NULL, 2}},
		},

		question897{
			para897{[]int{}},
			ans897{[]int{}},
		},

		question897{
			para897{[]int{1}},
			ans897{[]int{1}},
		},

		question897{
			para897{[]int{1, 2, 3}},
			ans897{[]int{2, NULL, 1, NULL, 3}},
		},

		question897{
			para897{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans897{[]int{3, NULL, 2, NULL, 4, NULL, 1, NULL, 4, NULL, 2, NULL, 3}},
		},

		question897{
			para897{[]int{1, 2, 2, NULL, 3, NULL, 3}},
			ans897{[]int{2, NULL, 3, NULL, 1, NULL, 2, NULL, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 897------------------------\n")

	for _, q := range qs {
		_, p := q.ans897, q.para897
		fmt.Printf("【input】:%v      ", p)
		rootOne := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", Tree2ints(increasingBST(rootOne)))
	}
	fmt.Printf("\n\n\n")
}
