package leetcode

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	r := 0
	w := 0
	b := 0 // label the end of different colors;
	for _, num := range nums {
		if num == 0 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		} else if num == 1 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		} else if num == 2 {
			b++
		}
	}
}
