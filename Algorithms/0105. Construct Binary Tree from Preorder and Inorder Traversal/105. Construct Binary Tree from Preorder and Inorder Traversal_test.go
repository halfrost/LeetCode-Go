package leetcode

import (
	"fmt"
	"testing"
)

type question105 struct {
	para105
	ans105
}

// para 是参数
// one 代表第一个参数
type para105 struct {
	preorder []int
	inorder  []int
}

// ans 是答案
// one 代表第一个答案
type ans105 struct {
	one []int
}

func Test_Problem105(t *testing.T) {

	qs := []question105{

		question105{
			para105{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			ans105{[]int{3, 9, 20, NULL, NULL, 15, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 105------------------------\n")

	for _, q := range qs {
		_, p := q.ans105, q.para105
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", Tree2ints(buildTree(p.preorder, p.inorder)))
	}
	fmt.Printf("\n\n\n")
}
