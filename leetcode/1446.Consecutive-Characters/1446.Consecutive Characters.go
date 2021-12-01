package leetcode

func maxPower(s string) int {
	cur := s[0]
	cnt := 1
	ans := 1
	for i := 1; i < len(s); i++ {
		if cur == s[i] {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt
			}
			cur = s[i]
			cnt = 1
		}
	}
	if cnt > ans {
		ans = cnt
	}
	return ans
}
