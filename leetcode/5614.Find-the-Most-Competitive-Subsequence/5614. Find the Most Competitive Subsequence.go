package leetcode

import (
	"fmt"
)

// 解法一 单调栈
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack)+len(nums)-i > k && len(stack) > 0 && nums[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}

// 解法二 DFS 超时
func mostCompetitive1(nums []int, k int) []int {
	c, visited, res := []int{}, map[int]bool{}, []int{}
	for i := 0; i < len(nums)-1; i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		} else {
			visited[nums[i]] = true
			generateIncSubsets(nums, i, k, c, &res)
		}
	}
	return res
}

func generateIncSubsets(nums []int, current, k int, c []int, res *[]int) {
	c = append(c, nums[current])
	fmt.Printf("c = %v res = %v\n", c, *res)
	if len(c) > k {
		return
	}
	if len(c) < k && len(*res) != 0 {
		b, flag := make([]int, len(c)), false
		copy(b, c)
		for i := 0; i < len(b); i++ {
			if b[i] < (*res)[i] {
				flag = true
				break
			}
		}
		if !flag {
			return
		}
	}
	// if len(*res) != 0 && len(c) <= len(*res) && c[len(c)-1] > (*res)[len(c)-1] {
	// 	return
	// }
	if len(c) == k {
		//fmt.Printf("c = %v\n", c)
		b, flag := make([]int, len(c)), false
		copy(b, c)
		if len(*res) == 0 {
			*res = b
		} else {
			for i := 0; i < len(b); i++ {
				if b[i] < (*res)[i] {
					flag = true
					break
				}
			}
			if flag {
				*res = b
			}
		}
		fmt.Printf("tmp = %v min = %v\n", b, *res)
	}
	visited := map[int]bool{}
	for i := current + 1; i < len(nums); i++ {
		// if nums[current] <= nums[i] {
		if _, ok := visited[nums[i]]; ok {
			continue
		} else {
			visited[nums[i]] = true
			generateIncSubsets(nums, i, k, c, res)
		}
		//}
	}
	c = c[:len(c)-1]
	return
}
