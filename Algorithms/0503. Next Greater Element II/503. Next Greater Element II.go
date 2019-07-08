package leetcode

// 解法一 单调栈
func nextGreaterElements(nums []int) []int {
	res := make([]int, 0)
	indexes := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, -1)
	}
	for i := 0; i < len(nums)*2; i++ {
		num := nums[i%len(nums)]
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] < num {
			index := indexes[len(indexes)-1]
			res[index] = num
			indexes = indexes[:len(indexes)-1]
		}
		indexes = append(indexes, i%len(nums))
	}
	return res
}

// 解法二
func nextGreaterElements1(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		j, find := (i+1)%len(nums), false
		for j != i {
			if nums[j] > nums[i] {
				find = true
				res = append(res, nums[j])
				break
			}
			j = (j + 1) % len(nums)
		}
		if !find {
			res = append(res, -1)
		}
	}
	return res
}
