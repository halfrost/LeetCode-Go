package leetcode

// 解法一 超简洁写法
func waysToMakeFair(nums []int) int {
	sum, res := [2]int{}, 0
	for i := 0; i < len(nums); i++ {
		sum[i%2] += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		sum[i%2] -= nums[i]
		if sum[i%2] == sum[1-(i%2)] {
			res++
		}
		sum[1-(i%2)] += nums[i]
	}
	return res
}

// 解法二 前缀和，后缀和
func waysToMakeFair1(nums []int) int {
	evenPrefix, oddPrefix, evenSuffix, oddSuffix, res := 0, 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenSuffix += nums[i]
		} else {
			oddSuffix += nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenSuffix -= nums[i]
		} else {
			oddSuffix -= nums[i]
		}
		if (evenPrefix + oddSuffix) == (oddPrefix + evenSuffix) {
			res++
		}
		if i%2 == 0 {
			evenPrefix += nums[i]
		} else {
			oddPrefix += nums[i]
		}
	}
	return res
}
