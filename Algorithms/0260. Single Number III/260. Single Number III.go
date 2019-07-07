package leetcode

func singleNumberIII(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// Get its last set bit (lsb)
	diff &= -diff
	res := []int{0, 0} // this array stores the two numbers we will return
	for _, num := range nums {
		if (num & diff) == 0 { // the bit is not set
			res[0] ^= num
		} else { // the bit is set
			res[1] ^= num
		}
	}
	return res
}
