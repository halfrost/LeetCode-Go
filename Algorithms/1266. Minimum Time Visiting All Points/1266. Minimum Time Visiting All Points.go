package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		res += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return res
}
