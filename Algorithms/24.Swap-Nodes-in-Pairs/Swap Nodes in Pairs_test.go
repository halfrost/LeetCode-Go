package leetcode

import (
	"fmt"
	"testing"
)

type question24 struct {
	para24
	ans24
}

// para 是参数
// one 代表第一个参数
type para24 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans24 struct {
	one []int
}

func Test_Problem24(t *testing.T) {

	qs := []question24{

		question24{
			para24{[]int{}},
			ans24{[]int{}},
		},

		question24{
			para24{[]int{1}},
			ans24{[]int{1}},
		},

		question24{
			para24{[]int{1, 2, 3, 4}},
			ans24{[]int{2, 1, 4, 3}},
		},

		question24{
			para24{[]int{1, 2, 3, 4, 5}},
			ans24{[]int{2, 1, 4, 3, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 24------------------------\n")

	for _, q := range qs {
		_, p := q.ans24, q.para24
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(swapPairs(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}

// convert *ListNode to []int
func L2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func S2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
