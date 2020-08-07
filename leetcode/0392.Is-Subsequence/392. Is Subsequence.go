package leetcode

// 解法一 O(n^2)
func isSubsequence(s string, t string) bool {
	index, flag := 0, false
	for i := 0; i < len(s); i++ {
		flag = false
		for ; index < len(t); index++ {
			if s[i] == t[index] {
				flag = true
				break
			}
		}
		if flag == true {
			index++
			continue
		} else {
			return false
		}
	}
	return true
}

// 解法二 O(n)
func isSubsequence1(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
