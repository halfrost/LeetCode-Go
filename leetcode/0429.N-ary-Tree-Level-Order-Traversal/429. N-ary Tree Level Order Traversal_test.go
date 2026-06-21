package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question429 struct {
	para429
	ans429
}

type para429 struct {
	root *Node
}

type ans429 struct {
	one [][]int
}

func Test_Problem429(t *testing.T) {
	// Tree:
	//        1
	//      / | \
	//     3  2  4
	//    / \
	//   5   6
	leaf5 := &Node{Val: 5}
	leaf6 := &Node{Val: 6}
	node3 := &Node{Val: 3, Children: []*Node{leaf5, leaf6}}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{node3, node2, node4}}

	qs := []question429{
		{
			para429{nil},
			ans429{nil},
		},
		{
			para429{&Node{Val: 1}},
			ans429{[][]int{{1}}},
		},
		{
			para429{root},
			ans429{[][]int{{1}, {3, 2, 4}, {5, 6}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 429------------------------\n")

	for _, q := range qs {
		a, p := q.ans429, q.para429
		got := levelOrder(p.root)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if !reflect.DeepEqual(got, a.one) {
			t.Fatalf("levelOrder(%v) = %v, want %v", p.root, got, a.one)
		}
	}

	fmt.Printf("\n\n\n")
}
