package leetcode

import "sort"

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	for i := 0; i < n; i++ {
		sort.Ints(cuboids[i]) // 立方体三边内部排序
	}
	// 立方体排序，先按最短边，再到最长边
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})
	res := 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
		res = max(res, dp[i])
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
