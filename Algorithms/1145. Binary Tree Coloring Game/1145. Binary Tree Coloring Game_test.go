package leetcode

import (
	"fmt"
	"testing"
)

type question1145 struct {
	para1145
	ans1145
}

// para 是参数
// one 代表第一个参数
type para1145 struct {
	root   []int
	target []int
	K      int
}

// ans 是答案
// one 代表第一个答案
type ans1145 struct {
	one []int
}

func Test_Problem1145(t *testing.T) {

	qs := []question1145{

		question1145{
			para1145{[]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}, []int{5}, 2},
			ans1145{[]int{7, 4, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1145------------------------\n")

	for _, q := range qs {
		_, p := q.ans1145, q.para1145
		tree, target := Ints2TreeNode(p.root), Ints2TreeNode(p.target)
		fmt.Printf("【input】:%v       【output】:%v\n", p, distanceK(tree, target, p.K))
	}
	fmt.Printf("\n\n\n")
}
