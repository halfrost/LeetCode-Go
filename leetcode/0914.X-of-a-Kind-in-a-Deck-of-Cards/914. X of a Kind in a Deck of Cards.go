package leetcode

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m, g := map[int]int{}, -1
	for _, d := range deck {
		m[d]++
	}
	for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
