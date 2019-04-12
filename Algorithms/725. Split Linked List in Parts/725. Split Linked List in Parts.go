package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, 0)
	if root == nil {
		for i := 0; i < k; i++ {
			res = append(res, nil)
		}
		return res
	}
	length := getLength(root)
	splitNum := length / k
	lengNum := length % k
	cur := root
	head := root
	pre := root
	fmt.Printf("总长度 %v, 分 %v 组, 前面 %v 组长度为 %v, 剩余 %v 组，每组 %v\n", length, k, lengNum, splitNum+1, k-lengNum, splitNum)
	if splitNum == 0 {
		for i := 0; i < k; i++ {
			if cur != nil {
				pre = cur.Next
				cur.Next = nil
				res = append(res, cur)
				cur = pre
			} else {
				res = append(res, nil)
			}
		}
		return res
	}
	for i := 0; i < lengNum; i++ {
		for j := 0; j < splitNum; j++ {
			cur = cur.Next
		}
		fmt.Printf("0 刚刚出来 head = %v cur = %v pre = %v\n", head, cur, head)
		pre = cur.Next
		cur.Next = nil
		res = append(res, head)
		head = pre
		cur = pre
		fmt.Printf("0 head = %v cur = %v pre = %v\n", head, cur, head)
	}
	for i := 0; i < k-lengNum; i++ {
		for j := 0; j < splitNum-1; j++ {
			cur = cur.Next
		}
		fmt.Printf("1 刚刚出来 head = %v cur = %v pre = %v\n", head, cur, head)
		pre = cur.Next
		cur.Next = nil
		res = append(res, head)
		head = pre
		cur = pre
	}
	return res
}
