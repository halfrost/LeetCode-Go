package leetcode

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	diff := 0
	maxDiff := 0
	for i, n2 := range nums2 {
		d := abs(nums1[i] - n2)
		diff += d
		if maxDiff < d {
			t := 100001
			for _, n1 := range nums1 {
				maxDiff = max(maxDiff, d-min(t, abs(n1-n2)))
			}
		}
	}
	return (diff - maxDiff) % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
