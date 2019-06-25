package leetcode

import "fmt"

func canPartition_(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	n, C, dp := len(nums), sum/2, make([]bool, sum/2+1)
	for i := 0; i <= C; i++ {
		dp[i] = (nums[0] == i)
	}
	fmt.Printf("dp = %v\n", dp)
	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			fmt.Printf("**dp = %v j = %v nums[i] = %v i = %v j-nums[i] = %v\n", dp, j, nums[i], i, j-nums[i])
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[C]
}

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum&1 == 1 {
		return false
	}

	half := sum / 2

	dp := make([]bool, half+1)
	dp[0] = true
	fmt.Printf("dp = %v\n", dp)
	for _, v := range nums {
		for j := half; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	fmt.Printf("*****dp = %v\n", dp)
	return dp[half]
}
