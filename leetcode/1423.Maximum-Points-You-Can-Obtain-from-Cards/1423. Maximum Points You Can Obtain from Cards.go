package leetcode

func maxScore(cardPoints []int, k int) int {
	windowSize, sum := len(cardPoints)-k, 0
	for _, val := range cardPoints[:windowSize] {
		sum += val
	}
	minSum := sum
	for i := windowSize; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-windowSize]
		if sum < minSum {
			minSum = sum
		}
	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}
