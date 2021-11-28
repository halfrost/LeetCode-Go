package leetcode

import "math"

func constructRectangle(area int) []int {
	ans := make([]int, 2)
	W := int(math.Sqrt(float64(area)))
	for W >= 1 {
		if area%W == 0 {
			ans[0], ans[1] = area/W, W
			break
		}
		W -= 1
	}
	return ans
}
