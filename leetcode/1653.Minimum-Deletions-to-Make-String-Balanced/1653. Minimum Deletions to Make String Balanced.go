package leetcode

// 解法一 DP
func minimumDeletions(s string) int {
	prev, res, bCount := 0, 0, 0
	for _, c := range s {
		if c == 'a' {
			res = min(prev+1, bCount)
			prev = res
		} else {
			bCount++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 解法二 模拟
func minimumDeletions1(s string) int {
	aCount, bCount, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount++
		}
	}
	res = aCount
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount--
		} else {
			bCount++
		}
		res = min(res, aCount+bCount)
	}
	return res
}
