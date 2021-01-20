package leetcode

func numIdenticalPairs(nums []int) int {
	total := 0
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x] == nums[y] {
				total++
			}
		}
	}
	return total
}
