package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1305 struct {
	para1305
	ans1305
}

// para 是参数
// one 代表第一个参数
type para1305 struct {
	root1 []int
	root2 []int
}

// ans 是答案
// one 代表第一个答案
type ans1305 struct {
	one []int
}

func Test_Problem1305(t *testing.T) {

	qs := []question1305{

		{
			para1305{[]int{2, 1, 4}, []int{1, 0, 3}},
			ans1305{[]int{0, 1, 1, 2, 3, 4}},
		},

		{
			para1305{[]int{0, -10, 10}, []int{5, 1, 7, 0, 2}},
			ans1305{[]int{-10, 0, 0, 1, 2, 5, 7, 10}},
		},

		{
			para1305{[]int{}, []int{5, 1, 7, 0, 2}},
			ans1305{[]int{0, 1, 2, 5, 7}},
		},

		{
			para1305{[]int{0, -10, 10}, []int{}},
			ans1305{[]int{-10, 0, 10}},
		},

		{
			para1305{[]int{1, structures.NULL, 8}, []int{8, 1}},
			ans1305{[]int{1, 1, 8, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1305------------------------\n")

	for _, q := range qs {
		_, p := q.ans1305, q.para1305
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", getAllElements(structures.Ints2TreeNode(p.root1), structures.Ints2TreeNode(p.root2)))
	}
	fmt.Printf("\n\n\n")
}
