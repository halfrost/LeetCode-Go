package leetcode

// 解法一 桶排序
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 || len(nums) < 2 {
		return false
	}
	buckets := map[int]int{}
	for i := 0; i < len(nums); i++ {
		// Get the ID of the bucket from element value nums[i] and bucket width t + 1
		key := nums[i] / (t + 1)
		// -7/9 = 0, but need -7/9 = -1
		if nums[i] < 0 {
			key--
		}
		if _, ok := buckets[key]; ok {
			return true
		}
		// check the lower bucket, and have to check the value too
		if v, ok := buckets[key-1]; ok && nums[i]-v <= t {
			return true
		}
		// check the upper bucket, and have to check the value too
		if v, ok := buckets[key+1]; ok && v-nums[i] <= t {
			return true
		}
		// maintain k size of window
		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1))
		}
		buckets[key] = nums[i]
	}
	return false
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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
