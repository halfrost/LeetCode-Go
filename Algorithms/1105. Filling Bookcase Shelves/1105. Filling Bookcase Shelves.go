package leetcode

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books)+1)
	dp[0] = 0
	for i := 1; i <= len(books); i++ {
		width, height := books[i-1][0], books[i-1][1]
		dp[i] = dp[i-1] + height
		for j := i - 1; j > 0 && width+books[j-1][0] <= shelfWidth; j-- {
			height = max(height, books[j-1][1])
			width += books[j-1][0]
			dp[i] = min(dp[i], dp[j-1]+height)
		}
	}
	return dp[len(books)]
}
