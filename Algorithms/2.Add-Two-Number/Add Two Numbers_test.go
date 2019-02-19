package leetcode

import (
	"fmt"
	"testing"
)

type question2 struct {
	para2
	ans2
}

// para 是参数
// one 代表第一个参数
type para2 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans2 struct {
	one []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{

		question2{
			para2{[]int{}, []int{}},
			ans2{[]int{}},
		},

		question2{
			para2{[]int{1}, []int{1}},
			ans2{[]int{2}},
		},

		question2{
			para2{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans2{[]int{2, 4, 6, 8}},
		},

		question2{
			para2{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans2{[]int{2, 4, 6, 8, 0, 1}},
		},

		question2{
			para2{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		question2{
			para2{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		question2{
			para2{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans2{[]int{7, 0, 8}},
		},

		question2{
			para2{[]int{1, 8, 3}, []int{7, 1}},
			ans2{[]int{8, 9, 3}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		_, p := q.ans2, q.para2
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(addTwoNumbers(S2l(p.one), S2l(p.another))))
	}
	fmt.Printf("\n\n\n")
}

// // convert *ListNode to []int
// func L2s(head *ListNode) []int {
// 	res := []int{}

// 	for head != nil {
// 		res = append(res, head.Val)
// 		head = head.Next
// 	}

// 	return res
// }

// // convert []int to *ListNode
// func S2l(nums []int) *ListNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	res := &ListNode{
// 		Val: nums[0],
// 	}
// 	temp := res
// 	for i := 1; i < len(nums); i++ {
// 		temp.Next = &ListNode{
// 			Val: nums[i],
// 		}
// 		temp = temp.Next
// 	}

// 	return res
// }
