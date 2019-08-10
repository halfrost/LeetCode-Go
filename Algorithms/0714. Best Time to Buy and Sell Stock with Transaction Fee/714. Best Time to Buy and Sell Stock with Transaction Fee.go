package leetcode

import (
	"math"
)

// 解法一 模拟 DP
func maxProfit714(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}
	return sell[len(sell)-1]
}

// 解法二 优化辅助空间的 DP
func maxProfit714_1(prices []int, fee int) int {
	sell, buy := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell
}
