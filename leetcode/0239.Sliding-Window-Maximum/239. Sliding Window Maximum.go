package leetcode

// 解法一 暴力解法 O(nk)
func maxSlidingWindow1(a []int, k int) []int {
	res := make([]int, 0, k)
	n := len(a)
	if n == 0 {
		return []int{}
	}
	for i := 0; i <= n-k; i++ {
		max := a[i]
		for j := 1; j < k; j++ {
			if max < a[i+j] {
				max = a[i+j]
			}
		}
		res = append(res, max)
	}

	return res
}

// 解法二 双端队列 Deque
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}
