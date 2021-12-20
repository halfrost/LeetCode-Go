package leetcode

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}
	judges := make(map[int]int)
	for _, v := range trust {
		judges[v[1]] += 1
	}
	for _, v := range trust {
		if _, ok := judges[v[0]]; ok {
			delete(judges, v[0])
		}
	}
	for k, v := range judges {
		if v == n-1 {
			return k
		}
	}
	return -1
}
