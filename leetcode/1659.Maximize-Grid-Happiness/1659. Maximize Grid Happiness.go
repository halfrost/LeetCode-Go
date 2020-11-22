package leetcode

import (
	"math"
)

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// lineStatus 					将每一行中 3 种状态进行编码，空白 - 0，内向人 - 1，外向人 - 2，每行状态用三进制表示
	// lineStatusList[729][6] 		每一行的三进制表示
	// introvertsCountInner[729] 	每一个 lineStatus 包含的内向人数
	// extrovertsCountInner[729] 	每一个 lineStatus 包含的外向人数
	// scoreInner[729] 			 	每一个 lineStatus 包含的行内得分（只统计 lineStatus 本身的得分，不包括它与上一行的）
	// scoreOuter[729][729] 	 	每一个 lineStatus 包含的行外得分
	// dp[上一行的 lineStatus][当前处理到的行][剩余的内向人数][剩余的外向人数]
	n3, lineStatus, introvertsCountInner, extrovertsCountInner, scoreInner, scoreOuter, lineStatusList, dp := math.Pow(3.0, float64(n)), 0, [729]int{}, [729]int{}, [729]int{}, [729][729]int{}, [729][6]int{}, [729][6][7][7]int{}
	for i := 0; i < 729; i++ {
		lineStatusList[i] = [6]int{}
	}
	for i := 0; i < 729; i++ {
		dp[i] = [6][7][7]int{}
		for j := 0; j < 6; j++ {
			dp[i][j] = [7][7]int{}
			for k := 0; k < 7; k++ {
				dp[i][j][k] = [7]int{-1, -1, -1, -1, -1, -1, -1}
			}
		}
	}
	// 预处理
	for lineStatus = 0; lineStatus < int(n3); lineStatus++ {
		tmp := lineStatus
		for i := 0; i < n; i++ {
			lineStatusList[lineStatus][i] = tmp % 3
			tmp /= 3
		}
		introvertsCountInner[lineStatus], extrovertsCountInner[lineStatus], scoreInner[lineStatus] = 0, 0, 0
		for i := 0; i < n; i++ {
			if lineStatusList[lineStatus][i] != 0 {
				// 个人分数
				if lineStatusList[lineStatus][i] == 1 {
					introvertsCountInner[lineStatus]++
					scoreInner[lineStatus] += 120
				} else if lineStatusList[lineStatus][i] == 2 {
					extrovertsCountInner[lineStatus]++
					scoreInner[lineStatus] += 40
				}
				// 行内分数
				if i-1 >= 0 {
					scoreInner[lineStatus] += closeScore(lineStatusList[lineStatus][i], lineStatusList[lineStatus][i-1])
				}
			}
		}
	}
	// 行外分数
	for lineStatus0 := 0; lineStatus0 < int(n3); lineStatus0++ {
		for lineStatus1 := 0; lineStatus1 < int(n3); lineStatus1++ {
			scoreOuter[lineStatus0][lineStatus1] = 0
			for i := 0; i < n; i++ {
				scoreOuter[lineStatus0][lineStatus1] += closeScore(lineStatusList[lineStatus0][i], lineStatusList[lineStatus1][i])
			}
		}
	}
	return dfs(0, 0, introvertsCount, extrovertsCount, m, int(n3), &dp, &introvertsCountInner, &extrovertsCountInner, &scoreInner, &scoreOuter)
}

// 如果 x 和 y 相邻，需要加上的分数
func closeScore(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	// 两个内向的人，每个人要 -30，一共 -60
	if x == 1 && y == 1 {
		return -60
	}
	if x == 2 && y == 2 {
		return 40
	}
	return -10
}

// dfs(上一行的 lineStatus，当前处理到的行，剩余的内向人数，剩余的外向人数）
func dfs(lineStatusLast, row, introvertsCount, extrovertsCount, m, n3 int, dp *[729][6][7][7]int, introvertsCountInner, extrovertsCountInner, scoreInner *[729]int, scoreOuter *[729][729]int) int {
	// 边界条件：如果已经处理完，或者没有人了
	if row == m || introvertsCount+extrovertsCount == 0 {
		return 0
	}
	// 记忆化
	if dp[lineStatusLast][row][introvertsCount][extrovertsCount] != -1 {
		return dp[lineStatusLast][row][introvertsCount][extrovertsCount]
	}
	best := 0
	for lineStatus := 0; lineStatus < n3; lineStatus++ {
		if introvertsCountInner[lineStatus] > introvertsCount || extrovertsCountInner[lineStatus] > extrovertsCount {
			continue
		}
		score := scoreInner[lineStatus] + scoreOuter[lineStatus][lineStatusLast]
		best = max(best, score+dfs(lineStatus, row+1, introvertsCount-introvertsCountInner[lineStatus], extrovertsCount-extrovertsCountInner[lineStatus], m, n3, dp, introvertsCountInner, extrovertsCountInner, scoreInner, scoreOuter))
	}
	dp[lineStatusLast][row][introvertsCount][extrovertsCount] = best
	return best
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
