package leetcode

import (
	"fmt"
	"testing"
)

type question559 struct {
	para559
	ans559
}

// para 是参数
type para559 struct {
	root *Node
}

// ans 是答案
type ans559 struct {
	ans int
}

func Test_Problem559(t *testing.T) {

	qs := []question559{

		{
			para559{&Node{
				Val: 1,
				Children: []*Node{
					{Val: 3, Children: []*Node{{Val: 5, Children: nil}, {Val: 6, Children: nil}}},
					{Val: 2, Children: nil},
					{Val: 4, Children: nil},
				},
			}},
			ans559{3},
		},

		{
			para559{&Node{
				Val: 1,
				Children: []*Node{
					{Val: 2, Children: nil},
					{Val: 3, Children: []*Node{{Val: 6, Children: nil}, {Val: 7, Children: []*Node{{Val: 11, Children: []*Node{{Val: 14, Children: nil}}}}}}},
					{Val: 4, Children: []*Node{{Val: 8, Children: []*Node{{Val: 12, Children: nil}}}}},
					{Val: 5, Children: []*Node{{Val: 9, Children: []*Node{{Val: 13, Children: nil}}}, {Val: 10, Children: nil}}},
				},
			}},
			ans559{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 559------------------------\n")

	for _, q := range qs {
		_, p := q.ans559, q.para559
		fmt.Printf("【input】:%v    【output】:%v\n", p, maxDepth(p.root))
	}
	fmt.Printf("\n\n\n")
}
