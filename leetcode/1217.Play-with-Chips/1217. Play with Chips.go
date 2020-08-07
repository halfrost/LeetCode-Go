package leetcode

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, c := range chips {
		if c%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
