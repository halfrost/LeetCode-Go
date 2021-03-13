package leetcode

import (
	"sort"
)

const mod = 1e9 + 7

// 解法一 DFS
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	numDict := map[int]bool{}
	for _, num := range arr {
		numDict[num] = true
	}
	dict, res := make(map[int][][2]int), 0
	for i, num := range arr {
		for j := i; j < len(arr) && num*arr[j] <= arr[len(arr)-1]; j++ {
			tmp := num * arr[j]
			if !numDict[tmp] {
				continue
			}
			dict[tmp] = append(dict[tmp], [2]int{num, arr[j]})
		}
	}
	cache := make(map[int]int)
	for _, num := range arr {
		res = (res + dfs(num, dict, cache)) % mod
	}
	return res
}

func dfs(num int, dict map[int][][2]int, cache map[int]int) int {
	if val, ok := cache[num]; ok {
		return val
	}
	res := 1
	for _, tuple := range dict[num] {
		a, b := tuple[0], tuple[1]
		x, y := dfs(a, dict, cache), dfs(b, dict, cache)
		tmp := x * y
		if a != b {
			tmp *= 2
		}
		res = (res + tmp) % mod
	}
	cache[num] = res
	return res
}

// 解法二 DP
func numFactoredBinaryTrees1(arr []int) int {
	dp := make(map[int]int)
	sort.Ints(arr)
	for i, curNum := range arr {
		for j := 0; j < i; j++ {
			factor := arr[j]
			quotient, remainder := curNum/factor, curNum%factor
			if remainder == 0 {
				dp[curNum] += dp[factor] * dp[quotient]
			}
		}
		dp[curNum]++
	}
	totalCount := 0
	for _, count := range dp {
		totalCount += count
	}
	return totalCount % mod
}
