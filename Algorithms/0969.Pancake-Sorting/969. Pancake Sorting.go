package leetcode

func pancakeSort(A []int) []int {
	if len(A) == 0 {
		return []int{}
	}
	right := len(A)
	var (
		ans []int
	)
	for right > 0 {
		idx := find(A, right)
		if idx != right-1 {
			reverse969(A, 0, idx)
			reverse969(A, 0, right-1)
			ans = append(ans, idx+1, right)
		}
		right--
	}

	return ans
}

func reverse969(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func find(nums []int, t int) int {
	for i, num := range nums {
		if num == t {
			return i
		}
	}
	return -1
}
