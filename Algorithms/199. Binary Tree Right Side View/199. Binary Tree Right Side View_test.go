package leetcode

import (
	"fmt"
	"testing"
)

type question199 struct {
	para199
	ans199
}

// para 是参数
// one 代表第一个参数
type para199 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans199 struct {
	one []int
}

func Test_Problem199(t *testing.T) {

	qs := []question199{

		question199{
			para199{[]int{}},
			ans199{[]int{}},
		},

		question199{
			para199{[]int{1}},
			ans199{[]int{1}},
		},

		question199{
			para199{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans199{[]int{3, 20, 7}},
		},

		question199{
			para199{[]int{1, 2, 3, 4, NULL, NULL, 5}},
			ans199{[]int{1, 3, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 199------------------------\n")

	for _, q := range qs {
		_, p := q.ans199, q.para199
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", rightSideView(root))
	}
	fmt.Printf("\n\n\n")
}
