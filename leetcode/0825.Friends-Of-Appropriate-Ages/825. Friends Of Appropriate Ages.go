package leetcocde

import "sort"

// 解法一 前缀和，时间复杂度 O(n)
func numFriendRequests(ages []int) int {
	count, prefixSum, res := make([]int, 121), make([]int, 121), 0
	for _, age := range ages {
		count[age]++
	}
	for i := 1; i < 121; i++ {
		prefixSum[i] = prefixSum[i-1] + count[i]
	}
	for i := 15; i < 121; i++ {
		if count[i] > 0 {
			bound := i/2 + 8
			res += count[i] * (prefixSum[i] - prefixSum[bound-1] - 1)
		}
	}
	return res
}

// 解法二 双指针 + 排序，时间复杂度 O(n logn)
func numFriendRequests1(ages []int) int {
	sort.Ints(ages)
	left, right, res := 0, 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		res += right - left
	}
	return res
}

// 解法三 暴力解法 O(n^2)
func numFriendRequests2(ages []int) int {
	res, count := 0, [125]int{}
	for _, x := range ages {
		count[x]++
	}
	for i := 1; i <= 120; i++ {
		for j := 1; j <= 120; j++ {
			if j > i {
				continue
			}
			if (j-7)*2 <= i {
				continue
			}
			if j > 100 && i < 100 {
				continue
			}
			if i != j {
				res += count[i] * count[j]
			} else {
				res += count[i] * (count[j] - 1)
			}
		}
	}
	return res
}
