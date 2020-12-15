package leetcode

func numWays(s string) int {
	ones := 0
	for _, c := range s {
		if c == '1' {
			ones++
		}
	}
	if ones%3 != 0 {
		return 0
	}
	if ones == 0 {
		return (len(s) - 1) * (len(s) - 2) / 2 % 1000000007
	}
	N, a, b, c, d, count := ones/3, 0, 0, 0, 0, 0
	for i, letter := range s {
		if letter == '0' {
			continue
		}
		if letter == '1' {
			count++
		}
		if count == N {
			a = i
		}
		if count == N+1 {
			b = i
		}
		if count == 2*N {
			c = i
		}
		if count == 2*N+1 {
			d = i
		}
	}
	return (b - a) * (d - c) % 1000000007
}
