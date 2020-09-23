package leetcode

func runningSum(nums []int) []int {
	result := []int{}
	counter := 0

	for x := 0; x < len(nums); x++ {
		for y := 0; y < x; y++ {
			counter += nums[y]
		}

		val := counter + nums[x]
		result = append(result, val)

		counter = 0
	}

	return result
}
