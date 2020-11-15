package leetcode

func minimumDeletions(s string) int {
	ai, bi, sum, temp, array := 0, 0, 0, 0, []int{}
	for ai = 0; ai < len(s); ai++ {
		if s[ai] == 'a' {
			break
		}
	}
	if ai != 0 && ai != len(s) {
		sum += ai
	}
	for bi = ai; bi < len(s); bi++ {
		if s[bi] == 'b' {
			break
		}
	}
	if s[bi-1] == 'a' {
		ai = bi - 1
	}
	if s[bi-1] == 'b' && bi != len(s) {
		ai = bi + 1
	}
	for j := bi; j < len(s); j++ {
		if s[j] == 'b' {
			temp++
		}
		if s[j] == 'a' && temp != 0 {
			array = append(array, temp)
			temp = 0
		}
	}
	if len(array) == 0 {
		return sum
	}
	dp := make([]int, len(array))
	dp[0] = min(array[0], len(array))
	for i := 1; i < len(array); i++ {
		dp[i] = min(dp[i-1]+array[i], dp[i-1]+len(array)-(i+1)+1)
	}
	return sum + dp[len(array)-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
