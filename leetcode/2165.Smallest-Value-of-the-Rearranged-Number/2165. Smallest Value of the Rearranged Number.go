package leetcode

import "sort"

func smallestNumber(num int64) int64 {
	pos := true
	if num < 0 {
		pos = false
		num *= -1
	}
	nums, m, res := []int{}, map[int]int{}, 0
	for num != 0 {
		tmp := int(num % 10)
		m[tmp]++
		num = num / 10
	}

	for k := range m {
		nums = append(nums, k)
	}
	if pos {
		sort.Ints(nums)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	}

	if nums[0] == 0 && len(nums) > 1 {
		res += nums[1]
		m[nums[1]]--
	}

	for _, v := range nums {
		if res != 0 {
			for j := m[v]; j > 0; j-- {
				res = res * 10
				res += v
			}
		} else {
			res += v
			tmp := m[v] - 1
			for j := tmp; j > 0; j-- {
				res = res * 10
				res += v
			}
		}
	}
	if !pos {
		return -1 * int64(res)
	}
	return int64(res)
}
