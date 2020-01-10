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
