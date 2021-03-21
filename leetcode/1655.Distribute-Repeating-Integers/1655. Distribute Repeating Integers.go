package leetcode

func canDistribute(nums []int, quantity []int) bool {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	return dfs(freq, quantity)
}

func dfs(freq map[int]int, quantity []int) bool {
	if len(quantity) == 0 {
		return true
	}
	visited := make(map[int]bool)
	for i := range freq {
		if visited[freq[i]] {
			continue
		}
		visited[freq[i]] = true
		if freq[i] >= quantity[0] {
			freq[i] -= quantity[0]
			if dfs(freq, quantity[1:]) {
				return true
			}
			freq[i] += quantity[0]
		}
	}
	return false
}
