package leetcode

import "sort"

// 解法一 排序 + 滑动窗口
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	elemList := make([]*elem, len(nums))
	for i, num := range nums {
		elemList[i] = &elem{num, i}
	}
	sort.SliceStable(elemList, func(i, j int) bool {
		if elemList[i].val != elemList[j].val {
			return elemList[i].val < elemList[j].val
		}
		return elemList[i].idx < elemList[j].idx
	})
	i, j := 0, 1
	for j < len(elemList) {
		if elemList[j].val-elemList[i].val <= t {
			if abs(elemList[j].idx-elemList[i].idx) <= k {
				return true
			}
			j++
		} else {
			i++
			if j <= i {
				j++
			}
		}
	}
	return false
}

type elem struct {
	val int
	idx int
}

// 解法二 滑动窗口 + 剪枝
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n && count < k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
			count++
		}
	}
	return false
}
