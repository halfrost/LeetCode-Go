package leetcode

func minimumDeviation(nums []int) int {
	min, max := 0, 0
	for i := range nums {
		if nums[i]%2 == 1 {
			nums[i] *= 2
		}
		if i == 0 {
			min = nums[i]
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		} else if max < nums[i] {
			max = nums[i]
		}
	}
	res := max - min
	for max%2 == 0 {
		tmax, tmin := 0, 0
		for i := range nums {
			if nums[i] == max || (nums[i]%2 == 0 && min <= nums[i]/2) {
				nums[i] /= 2
			}
			if i == 0 {
				tmin = nums[i]
				tmax = nums[i]
			} else if nums[i] < tmin {
				tmin = nums[i]
			} else if tmax < nums[i] {
				tmax = nums[i]
			}
		}
		if tmax-tmin < res {
			res = tmax - tmin
		}
		min, max = tmin, tmax
	}
	return res
}
