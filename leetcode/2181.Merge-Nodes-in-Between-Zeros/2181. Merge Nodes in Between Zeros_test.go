package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question2181 struct {
	para2181
	ans2181
}

// para 是参数
// one 代表第一个参数
type para2181 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans2181 struct {
	one []int
}

func Test_Problem2181(t *testing.T) {

	qs := []question2181{

		{
			para2181{[]int{0, 3, 1, 0, 4, 5, 2, 0}},
			ans2181{[]int{4, 11}},
		},

		{
			para2181{[]int{0, 1, 0, 3, 0, 2, 2, 0}},
			ans2181{[]int{1, 3, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2181------------------------\n")

	for _, q := range qs {
		_, p := q.ans2181, q.para2181
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(mergeNodes(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return newHead.Next
}
