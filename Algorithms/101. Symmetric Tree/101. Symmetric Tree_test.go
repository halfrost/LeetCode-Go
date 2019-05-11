package leetcode

import (
	"fmt"
	"testing"
)

type question101 struct {
	para101
	ans101
}

// para 是参数
// one 代表第一个参数
type para101 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans101 struct {
	one bool
}

func Test_Problem101(t *testing.T) {

	qs := []question101{

		question101{
			para101{[]int{3, 4, 4, 5, NULL, NULL, 5, 6, NULL, NULL, 6}},
			ans101{true},
		},

		question101{
			para101{[]int{1, 2, 2, NULL, 3, 3}},
			ans101{true},
		},

		question101{
			para101{[]int{}},
			ans101{true},
		},

		question101{
			para101{[]int{1}},
			ans101{true},
		},

		question101{
			para101{[]int{1, 2, 3}},
			ans101{false},
		},

		question101{
			para101{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans101{true},
		},

		question101{
			para101{[]int{1, 2, 2, NULL, 3, NULL, 3}},
			ans101{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 101------------------------\n")

	for _, q := range qs {
		_, p := q.ans101, q.para101
		fmt.Printf("【input】:%v      ", p)
		rootOne := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isSymmetric(rootOne))
	}
	fmt.Printf("\n\n\n")
}
