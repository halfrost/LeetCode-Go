package leetcode

import (
	"fmt"
	"testing"
)

type question1609 struct {
	para1609
	ans1609
}

// para 是参数
type para1609 struct {
	root *TreeNode
}

// ans 是答案
type ans1609 struct {
	ans bool
}

func Test_Problem1609(t *testing.T) {

	qs := []question1609{

		{
			para1609{&TreeNode{1,
				&TreeNode{10,
					&TreeNode{3, &TreeNode{12, nil, nil}, &TreeNode{8, nil, nil}}, nil},
				&TreeNode{4, &TreeNode{7, &TreeNode{6, nil, nil}, nil}, &TreeNode{9, nil, &TreeNode{2, nil, nil}}}}},
			ans1609{true},
		},

		{
			para1609{&TreeNode{5,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{3, nil, nil}},
				&TreeNode{2, &TreeNode{7, nil, nil}, nil}}},
			ans1609{false},
		},

		{
			para1609{&TreeNode{5,
				&TreeNode{9, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{1, &TreeNode{7, nil, nil}, nil}}},
			ans1609{false},
		},

		{
			para1609{&TreeNode{1, nil, nil}},
			ans1609{true},
		},

		{
			para1609{&TreeNode{11,
				&TreeNode{8,
					&TreeNode{1, &TreeNode{30, &TreeNode{17, nil, nil}, nil}, &TreeNode{20, nil, nil}},
					&TreeNode{3, &TreeNode{18, nil, nil}, &TreeNode{16, nil, nil}}},
				&TreeNode{6,
					&TreeNode{9, &TreeNode{12, nil, nil}, &TreeNode{10, nil, nil}},
					&TreeNode{11, &TreeNode{4, nil, nil}, &TreeNode{2, nil, nil}}}}},
			ans1609{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1609------------------------\n")

	for _, q := range qs {
		_, p := q.ans1609, q.para1609
		fmt.Printf("【input】:%v    【output】:%v\n", p.root, isEvenOddTree(p.root))
	}
	fmt.Printf("\n\n\n")
}
