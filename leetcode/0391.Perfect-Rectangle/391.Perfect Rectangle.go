package leetcode

type point struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {
	minX, minY, maxA, maxB := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	area := 0
	cnt := make(map[point]int)
	for _, v := range rectangles {
		x, y, a, b := v[0], v[1], v[2], v[3]
		area += (a - x) * (b - y)
		minX = min(minX, x)
		minY = min(minY, y)
		maxA = max(maxA, a)
		maxB = max(maxB, b)
		cnt[point{x, y}]++
		cnt[point{a, b}]++
		cnt[point{x, b}]++
		cnt[point{a, y}]++
	}
	if area != (maxA-minX)*(maxB-minY) ||
		cnt[point{minX, minY}] != 1 || cnt[point{maxA, maxB}] != 1 ||
		cnt[point{minX, maxB}] != 1 || cnt[point{maxA, minY}] != 1 {
		return false
	}
	delete(cnt, point{minX, minY})
	delete(cnt, point{maxA, maxB})
	delete(cnt, point{minX, maxB})
	delete(cnt, point{maxA, minY})
	for _, v := range cnt {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
