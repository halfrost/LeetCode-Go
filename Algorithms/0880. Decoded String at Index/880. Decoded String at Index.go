package leetcode

func isLetter(char byte) bool {
	if char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

func decodeAtIndex(S string, K int) string {
	length := 0
	for i := 0; i < len(S); i++ {
		if isLetter(S[i]) {
			length++
			if length == K {
				return string(S[i])
			}
		} else {
			if length*int(S[i]-'0') >= K {
				if K%length != 0 {
					return decodeAtIndex(S[:i], K%length)
				}
				return decodeAtIndex(S[:i], length)
			}
			length *= int(S[i] - '0')
		}
	}
	return ""
}
