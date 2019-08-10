package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法一 单调栈
func nextLargerNodes(head *ListNode) []int {
	res, indexes, nums := make([]int, 0), make([]int, 0), make([]int, 0)
	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	for i := 0; i < len(nums); i++ {
		res = append(res, 0)
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] < num {
			index := indexes[len(indexes)-1]
			res[index] = num
			indexes = indexes[:len(indexes)-1]
		}
		indexes = append(indexes, i)
	}
	return res
}
