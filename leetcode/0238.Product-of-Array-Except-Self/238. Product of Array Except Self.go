package leetcode

// 解法一 前缀积 * 后缀积。不使用除法，时间复杂度 O(n)，除输出数组外空间复杂度 O(1)。
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	// res[i] 先存 nums[i] 左边所有元素的乘积
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	// 再从右往左乘上 nums[i] 右边所有元素的乘积，right 累计后缀积
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
