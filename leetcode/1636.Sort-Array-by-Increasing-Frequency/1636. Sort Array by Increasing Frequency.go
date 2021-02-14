package leetcode

import "sort"

func frequencySort(nums []int) []int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[j] < nums[i]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}
