package leetcode

func largestAltitude(gain []int) int {
	max, height := 0, 0
	for _, g := range gain {
		height += g
		if height > max {
			max = height
		}
	}
	return max
}
