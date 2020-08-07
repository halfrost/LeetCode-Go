package leetcode

func numRabbits(ans []int) int {
	total, m := 0, make(map[int]int)
	for _, v := range ans {
		if m[v] == 0 {
			m[v] += v
			total += v + 1
		} else {
			m[v]--
		}
	}
	return total
}
