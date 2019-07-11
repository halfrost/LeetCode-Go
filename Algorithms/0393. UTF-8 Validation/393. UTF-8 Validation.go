package leetcode

func validUtf8(data []int) bool {
	count := 0
	for _, d := range data {
		if count == 0 {
			if d >= 248 { // 11111000 = 248
				return false
			} else if d >= 240 { // 11110000 = 240
				count = 3
			} else if d >= 224 { // 11100000 = 224
				count = 2
			} else if d >= 192 { // 11000000 = 192
				count = 1
			} else if d > 127 { // 01111111 = 127
				return false
			}
		} else {
			if d <= 127 || d >= 192 {
				return false
			}
			count--
		}
	}
	return count == 0
}
