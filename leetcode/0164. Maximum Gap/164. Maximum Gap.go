package leetcode

// 解法一
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	quickSort164(nums, 0, len(nums)-1)

	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i+1] - nums[i]) > res {
			res = nums[i+1] - nums[i]
		}
	}
	return res
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}

// 解法二
func maximumGap1(nums []int) int {

	if nums == nil || len(nums) < 2 {
		return 0
	}

	// m is the maximal number in nums
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
	}

	exp := 1 // 1, 10, 100, 1000 ...
	R := 10  // 10 digits

	aux := make([]int, len(nums))

	for (m / exp) > 0 { // Go through all digits from LSB to MSB
		count := make([]int, R)

		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}

		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}

		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
		}

		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}

	maxValue := 0
	for i := 1; i < len(aux); i++ {
		maxValue = max(maxValue, aux[i]-aux[i-1])
	}

	return maxValue
}
