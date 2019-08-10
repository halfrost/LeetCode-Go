package leetcode

func totalHammingDistance(nums []int) int {
	total, n := 0, len(nums)
	for i := 0; i < 32; i++ {
		bitCount := 0
		for j := 0; j < n; j++ {
			bitCount += (nums[j] >> uint(i)) & 1
		}
		total += bitCount * (n - bitCount)
	}
	return total
}

// 暴力解法超时！
func totalHammingDistance1(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += hammingDistance(nums[i], nums[j])
		}
	}
	return res
}
