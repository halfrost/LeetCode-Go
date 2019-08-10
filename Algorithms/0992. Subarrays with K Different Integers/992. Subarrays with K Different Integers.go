package leetcode

func subarraysWithKDistinct(A []int, K int) int {
	return subarraysWithKDistinctSlideWindow(A, K) - subarraysWithKDistinctSlideWindow(A, K-1)
}

func subarraysWithKDistinctSlideWindow(A []int, K int) int {
	left, right, counter, res, freq := 0, 0, K, 0, map[int]int{}
	for right = 0; right < len(A); right++ {
		if freq[A[right]] == 0 {
			counter--
		}
		freq[A[right]]++
		for counter < 0 {
			freq[A[left]]--
			if freq[A[left]] == 0 {
				counter++
			}
			left++
		}
		res += right - left + 1
	}
	return res
}
