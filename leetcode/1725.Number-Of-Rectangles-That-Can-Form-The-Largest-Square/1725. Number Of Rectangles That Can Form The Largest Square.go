package leetcode

func countGoodRectangles(rectangles [][]int) int {
	minLength, count := 0, 0
	for i, _ := range rectangles {
		minSide := 0
		if rectangles[i][0] <= rectangles[i][1] {
			minSide = rectangles[i][0]
		} else {
			minSide = rectangles[i][1]
		}
		if minSide > minLength {
			minLength = minSide
			count = 1
		} else if minSide == minLength {
			count++
		}
	}
	return count
}
