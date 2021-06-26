package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome234(head *ListNode) bool {
	slice := []int{}

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	for i, j := 0, len(slice)-1; i < j; {
		if slice[i] != slice[j] {
			return false
		}

		i++
		j--
	}

	return true
}
