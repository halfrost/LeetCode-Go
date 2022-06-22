package leetcode

func modifyString(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for b := byte('a'); b <= 'z'; b++ {
				if !(i > 0 && res[i-1] == b || i < len(res)-1 && res[i+1] == b) {
					res[i] = b
					break
				}
			}
		}
	}
	return string(res)
}
