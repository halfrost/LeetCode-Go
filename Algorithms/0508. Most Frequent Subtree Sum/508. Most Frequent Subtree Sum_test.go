package leetcode

import (
	"fmt"
	"testing"
)

type question508 struct {
	para508
	ans508
}

// para 是参数
// one 代表第一个参数
type para508 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans508 struct {
	one []int
}

func Test_Problem508(t *testing.T) {

	qs := []question508{

		question508{
			para508{[]int{}},
			ans508{[]int{}},
		},

		question508{
			para508{[]int{1, 1}},
			ans508{[]int{1, 2}},
		},

		question508{
			para508{[]int{1}},
			ans508{[]int{1}},
		},

		question508{
			para508{[]int{5, 2, -3}},
			ans508{[]int{2, -3, 4}},
		},

		question508{
			para508{[]int{5, 2, -5}},
			ans508{[]int{2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 508------------------------\n")

	for _, q := range qs {
		_, p := q.ans508, q.para508
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", findFrequentTreeSum(root))
	}
	fmt.Printf("\n\n\n")
}
