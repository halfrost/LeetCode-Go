package leetcode

func uniqueLetterString(S string) int {
	res, left, right := 0, 0, 0
	for i := 0; i < len(S); i++ {
		left = i - 1
		for left >= 0 && S[left] != S[i] {
			left--
		}
		right = i + 1
		for right < len(S) && S[right] != S[i] {
			right++
		}
		res += (i - left) * (right - i)
	}
	return res % 1000000007
}

// 暴力解法，超时！时间复杂度 O(n^2)
func uniqueLetterString1(S string) int {
	if len(S) == 0 {
		return 0
	}
	letterMap := map[byte]int{}
	res, mod := 0, 1000000007
	for i := 0; i < len(S); i++ {
		letterMap = map[byte]int{}
		for j := i; j < len(S); j++ {
			letterMap[S[j]]++
			tmp := 0
			for _, v := range letterMap {
				if v > 1 {
					tmp++
				}
			}
			if tmp == len(letterMap) {
				continue
			} else {
				res += len(letterMap) - tmp
			}
		}
	}
	return res % mod
}
