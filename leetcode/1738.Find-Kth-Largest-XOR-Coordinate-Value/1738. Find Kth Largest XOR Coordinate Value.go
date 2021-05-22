package leetcode

import "sort"

// 解法一 压缩版的前缀和
func kthLargestValue(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res, prefixSum := make([]int, 0, len(matrix)*len(matrix[0])), make([]int, len(matrix[0]))
	for i := range matrix {
		line := 0
		for j, v := range matrix[i] {
			line ^= v
			prefixSum[j] ^= line
			res = append(res, prefixSum[j])
		}
	}
	sort.Ints(res)
	return res[len(res)-k]
}

// 解法二 前缀和
func kthLargestValue1(matrix [][]int, k int) int {
	nums, prefixSum := []int{}, make([][]int, len(matrix)+1)
	prefixSum[0] = make([]int, len(matrix[0])+1)
	for i, row := range matrix {
		prefixSum[i+1] = make([]int, len(matrix[0])+1)
		for j, val := range row {
			prefixSum[i+1][j+1] = prefixSum[i+1][j] ^ prefixSum[i][j+1] ^ prefixSum[i][j] ^ val
			nums = append(nums, prefixSum[i+1][j+1])
		}
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}
