package leetcode

func balancedString(s string) int {
	count, k := make([]int, 128), len(s)/4
	for _, v := range s {
		count[int(v)]++
	}
	left, right, res := 0, -1, len(s)
	for left < len(s) {
		if count['Q'] > k || count['W'] > k || count['E'] > k || count['R'] > k {
			if right+1 < len(s) {
				right++
				count[s[right]]--
			} else {
				break
			}
		} else {
			res = min(res, right-left+1)
			count[s[left]]++
			left++
		}
	}
	return res
}
