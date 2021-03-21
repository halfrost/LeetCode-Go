package leetcode

//解法一 优化版 prefixSum + sufixSum
func getSumAbsoluteDifferences(nums []int) []int {
	size := len(nums)
	sufixSum := make([]int, size)
	sufixSum[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		sufixSum[i] = sufixSum[i+1] + nums[i]
	}
	ans, preSum := make([]int, size), 0
	for i := 0; i < size; i++ {
		// 后面可以加到的值
		res, sum := 0, sufixSum[i]-nums[i]
		res += (sum - (size-i-1)*nums[i])
		// 前面可以加到的值
		res += (i*nums[i] - preSum)
		ans[i] = res
		preSum += nums[i]
	}
	return ans
}

// 解法二 prefixSum
func getSumAbsoluteDifferences1(nums []int) []int {
	preSum, res, sum := []int{}, []int{}, nums[0]
	preSum = append(preSum, nums[0])
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		preSum = append(preSum, sum)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res = append(res, preSum[len(nums)-1]-preSum[0]-nums[i]*(len(nums)-1))
		} else if i > 0 && i < len(nums)-1 {
			res = append(res, preSum[len(nums)-1]-preSum[i]-preSum[i-1]+nums[i]*i-nums[i]*(len(nums)-1-i))
		} else {
			res = append(res, nums[i]*len(nums)-preSum[len(nums)-1])
		}
	}
	return res
}
