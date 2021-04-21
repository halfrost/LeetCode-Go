package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question589 struct {
	para589
	ans589
}

// para 是参数
// one 代表第一个参数
type para589 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans589 struct {
	one []int
}

func Test_Problem589(t *testing.T) {

	qs := []question589{

		{
			para589{[]int{1, structures.NULL, 3, 2, 4, structures.NULL, 5, 6}},
			ans589{[]int{1, 3, 5, 6, 2, 4}},
		},

		{
			para589{[]int{1, structures.NULL, 2, 3, 4, 5, structures.NULL, structures.NULL, 6, 7, structures.NULL, 8, structures.NULL, 9, 10, structures.NULL, structures.NULL, 11, structures.NULL, 12, structures.NULL, 13, structures.NULL, structures.NULL, 14}},
			ans589{[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 589------------------------\n")

	for _, q := range qs {
		_, p := q.ans589, q.para589
		fmt.Printf("【input】:%v      ", p)
		rootOne := int2NaryNode(p.one)
		fmt.Printf("【output】:%v      \n", preorder(rootOne))
	}
	fmt.Printf("\n\n\n")
}

func int2NaryNode(nodes []int) *Node {
	root := &Node{}
	if len(nodes) > 1 {
		root.Val = nodes[0]
	}
	queue := []*Node{}
	queue = append(queue, root)
	i := 1
	count := 0
	for i < len(nodes) {
		node := queue[0]

		childrens := []*Node{}
		for ; i < len(nodes) && nodes[i] != structures.NULL; i++ {
			tmp := &Node{Val: nodes[i]}
			childrens = append(childrens, tmp)
			queue = append(queue, tmp)
		}
		count++
		if count%2 == 0 {
			queue = queue[1:]
			count = 1
		}
		if node != nil {
			node.Children = childrens
		}
		i++
	}
	return root
}
