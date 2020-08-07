package leetcode

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

// 以下是拓展题
// 在数组中每个元素都出现 5 次，找出只出现 1 次的数。

// 解法一
func singleNumberIIIII(nums []int) int {
	na, nb, nc := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nb = nb ^ (nums[i] & na)
		na = (na ^ nums[i]) & ^nc
		nc = nc ^ (nums[i] & ^na & ^nb)
	}
	return na & ^nb & ^nc
}

// 解法二
func singleNumberIIIII1(nums []int) int {
	twos, threes, ones := 0xffffffff, 0xffffffff, 0
	for i := 0; i < len(nums); i++ {
		threes = threes ^ (nums[i] & twos)
		twos = (twos ^ nums[i]) & ^ones
		ones = ones ^ (nums[i] & ^twos & ^threes)
	}
	return ones
}
