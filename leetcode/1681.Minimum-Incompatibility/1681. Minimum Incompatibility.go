package leetcode

import (
	"math"
	"sort"
)

func minimumIncompatibility(nums []int, k int) int {
	sort.Ints(nums)
	eachSize, counts := len(nums)/k, make([]int, len(nums)+1)
	for i := range nums {
		counts[nums[i]]++
		if counts[nums[i]] > k {
			return -1
		}
	}
	orders := []int{}
	for i := range counts {
		orders = append(orders, i)
	}
	sort.Ints(orders)
	res := math.MaxInt32
	generatePermutation1681(nums, counts, orders, 0, 0, eachSize, &res, []int{})
	// 当所有数字出现次数均不超过 k 时，必然存在一个合法划分，
	// 因此 res 一定会被更新，无需再判断 res == math.MaxInt32。
	return res
}

func generatePermutation1681(nums, counts, order []int, index, sum, eachSize int, res *int, current []int) {
	if len(current) > 0 && len(current)%eachSize == 0 {
		sum += current[len(current)-1] - current[len(current)-eachSize]
		index = 0
	}
	if sum >= *res {
		return
	}
	if len(current) == len(nums) {
		if sum < *res {
			*res = sum
		}
		return
	}
	for i := index; i < len(counts); i++ {
		if counts[order[i]] == 0 {
			continue
		}
		counts[order[i]]--
		current = append(current, order[i])
		generatePermutation1681(nums, counts, order, i+1, sum, eachSize, res, current)
		current = current[:len(current)-1]
		counts[order[i]]++
		// 这里是关键的剪枝
		if index == 0 {
			break
		}
	}
}
