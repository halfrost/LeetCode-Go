package leetcode

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			// pop stack
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
		}
		// push stack
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
