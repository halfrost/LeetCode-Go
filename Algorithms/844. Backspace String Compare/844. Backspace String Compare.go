package leetcode

func backspaceCompare(S string, T string) bool {
	s := make([]rune, 0)
	for _, c := range S {
		if c == '#' {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, c)
		}
	}
	s2 := make([]rune, 0)
	for _, c := range T {
		if c == '#' {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		} else {
			s2 = append(s2, c)
		}
	}
	return string(s) == string(s2)
}
