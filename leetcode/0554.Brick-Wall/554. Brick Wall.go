package leetcode

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	for _, row := range wall {
		sum := 0
		for i := 0; i < len(row)-1; i++ {
			sum += row[i]
			m[sum]++
		}
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}
