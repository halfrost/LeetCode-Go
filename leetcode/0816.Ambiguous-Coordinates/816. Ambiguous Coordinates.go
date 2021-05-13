package leetcode

func ambiguousCoordinates(s string) []string {
	res := []string{}
	s = s[1 : len(s)-1]
	for i := range s[:len(s)-1] {
		a := build(s[:i+1])
		b := build(s[i+1:])
		for _, ta := range a {
			for _, tb := range b {
				res = append(res, "("+ta+", "+tb+")")
			}
		}
	}
	return res
}

func build(s string) []string {
	res := []string{}
	if len(s) == 1 || s[0] != '0' {
		res = append(res, s)
	}
	// 结尾带 0 的情况
	if s[len(s)-1] == '0' {
		return res
	}
	// 切分长度大于一位且带前导 0 的情况
	if s[0] == '0' {
		res = append(res, "0."+s[1:])
		return res
	}
	for i := range s[:len(s)-1] {
		res = append(res, s[:i+1]+"."+s[i+1:])
	}
	return res
}
