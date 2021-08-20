package leetcode

import "sort"

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	total := 0
	for _, v := range matchsticks {
		total += v
	}
	if total%4 != 0 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	visited := make([]bool, 16)
	return dfs(matchsticks, 0, 0, 0, total, &visited)
}

func dfs(matchsticks []int, cur, group, sum, total int, visited *[]bool) bool {
	if group == 4 {
		return true
	}
	if sum > total/4 {
		return false
	}
	if sum == total/4 {
		return dfs(matchsticks, 0, group+1, 0, total, visited)
	}
	last := -1
	for i := cur; i < len(matchsticks); i++ {
		if (*visited)[i] {
			continue
		}
		if last == matchsticks[i] {
			continue
		}
		(*visited)[i] = true
		last = matchsticks[i]
		if dfs(matchsticks, i+1, group, sum+matchsticks[i], total, visited) {
			return true
		}
		(*visited)[i] = false
	}
	return false
}
