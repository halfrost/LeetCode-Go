package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	level := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var nums []int
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			nums = append(nums, node.Val)
		}
		if level%2 == 0 {
			if !even(nums) {
				return false
			}
		} else {
			if !odd(nums) {
				return false
			}
		}
		queue = queue[length:]
		level++
	}
	return true
}

func odd(nums []int) bool {
	cur := nums[0]
	if cur%2 != 0 {
		return false
	}
	for _, num := range nums[1:] {
		if num >= cur {
			return false
		}
		if num%2 != 0 {
			return false
		}
		cur = num
	}
	return true
}

func even(nums []int) bool {
	cur := nums[0]
	if cur%2 == 0 {
		return false
	}
	for _, num := range nums[1:] {
		if num <= cur {
			return false
		}
		if num%2 == 0 {
			return false
		}
		cur = num
	}
	return true
}
