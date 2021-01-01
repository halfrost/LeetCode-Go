package leetcode

func reformatNumber(number string) string {
	str, count := "", 0
	for i := 0; i < len(number); i++ {
		if number[i] != '-' && number[i] != ' ' {
			str += string(number[i])
			// str = append(str, number[i])
		}
	}
	if len(str) == 4 {
		str = str[:2] + "-" + str[2:]
		return str
	}
	if len(str) > 3 {
		if (len(str)-4)%3 == 0 {
			for i := len(str) - 5; i >= 0; i-- {
				count++
				if count%3 == 0 && i != 0 {
					str = str[:i] + "-" + str[i:]
				}
			}
			str = str[:len(str)-2] + "-" + str[len(str)-2:]
			str = str[:len(str)-5] + "-" + str[len(str)-5:]
			return str
		}
		if (len(str)-2)%3 == 0 {
			for i := len(str) - 3; i >= 0; i-- {
				count++
				if count%3 == 0 && i != 0 {
					str = str[:i] + "-" + str[i:]
				}
			}
			str = str[:len(str)-2] + "-" + str[len(str)-2:]
			return str
		}
		length := len(str)
		count = 0
		for j := 0; j < len(str); j++ {
			count++
			if count%3 == 0 && count != length {
				// head :=
				// tail :=
				str = str[:j+1] + "-" + str[j+1:]
				j++
			}
		}

	}
	return str
}
