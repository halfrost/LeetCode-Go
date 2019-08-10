package leetcode

import "sort"

// 解法一
func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

// 解法二
func subsets1(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, org := range res {
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}
