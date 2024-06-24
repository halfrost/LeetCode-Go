package leetcode

// 解法一 DP
func findTargetSumWays(nums []int, S int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	if S > total || (S+total)%2 == 1 || S+total < 0 {
		return 0
	}
	target := (S + total) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

// 解法二 DFS
func findTargetSumWays1(nums []int, S int) int {
	// sums[i] 存储的是后缀和 nums[i:]，即从 i 到结尾的和
	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	res := 0
	dfsFindTargetSumWays(nums, 0, 0, S, &res, sums)
	return res
}

func dfsFindTargetSumWays(nums []int, index int, curSum int, S int, res *int, sums []int) {
	if index == len(nums) {
		if curSum == S {
			*(res) = *(res) + 1
		}
		return
	}
	// 剪枝优化：如果 sums[index] 值小于剩下需要正数的值，那么右边就算都是 + 号也无能为力了，所以这里可以剪枝了
	if S-curSum > sums[index] {
		return
	}
	dfsFindTargetSumWays(nums, index+1, curSum+nums[index], S, res, sums)
	dfsFindTargetSumWays(nums, index+1, curSum-nums[index], S, res, sums)
}

// recursive 
func RecursivefindTargetSumWays(nums []int, target int) int {
	//  Runtime: 611 ms, faster than 19.64% of Go online submissions for Target Sum.
	//  Memory Usage: 2 MB, less than 97.26% of Go online submissions for Target Sum.
	ans :=0
	CurrentValue:=0 
	var sum func(int,[]int)
	sum =func (CurrentValue int, arr []int) {
		if len(arr)==0 && CurrentValue==target {
			ans+=1
			return
		}else if len(arr)==0 {
			return
		}
		temp:=arr
		if len(arr)==1{
			temp=[]int{}
		}else {
			temp=arr[1:]
		}

		sum(CurrentValue+arr[0],temp)
		sum(CurrentValue-arr[0],temp)


	}

	sum(CurrentValue,nums)
	return ans 
}

