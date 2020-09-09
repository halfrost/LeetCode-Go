package leetcode

func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 && len(typed) == 0 {
		return true
	}
	if (len(name) == 0 && len(typed) != 0) || (len(name) != 0 && len(typed) == 0) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		for i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		}
		for j < len(typed) && typed[j] == typed[j-1] {
			j++
		}
	}
	return i == len(name) && j == len(typed)
}
