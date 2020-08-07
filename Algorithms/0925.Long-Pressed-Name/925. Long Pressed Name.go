package leetcode

func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 && len(typed) == 0 {
		return true
	}
	if (len(name) == 0 && len(typed) != 0) || (len(name) != 0 && len(typed) == 0) {
		return false
	}

	j := 0
	for i := 0; i < len(name); i++ {
		if j < len(typed) && name[i] == typed[j] {
			j++
			continue
		} else {
			if i > 0 && j < len(typed) && name[i-1] == typed[j] {
				j++
				i--
			} else {
				return false
			}
		}
	}
	return true
}
