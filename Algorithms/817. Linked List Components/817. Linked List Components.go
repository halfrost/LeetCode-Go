package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func numComponents(head *ListNode, G []int) int {
	if head.Next == nil {
		return 1
	}
	gMap := toMap(G)
	count := 0
	cur := head

	for cur != nil {
		if _, ok := gMap[cur.Val]; ok {
			if cur.Next == nil { // 末尾存在，直接加一
				count++
			} else {
				if _, ok = gMap[cur.Next.Val]; !ok {
					count++
				}
			}
		}
		cur = cur.Next
	}
	return count
}

func toMap(G []int) map[int]int {
	GMap := make(map[int]int, 0)
	for _, value := range G {
		GMap[value] = 0
	}
	return GMap
}
