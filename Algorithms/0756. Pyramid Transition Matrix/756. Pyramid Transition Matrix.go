package leetcode

func pyramidTransition(bottom string, allowed []string) bool {
	pyramid := make(map[string][]string)
	for _, v := range allowed {
		pyramid[v[:len(v)-1]] = append(pyramid[v[:len(v)-1]], string(v[len(v)-1]))
	}
	return dfsT(bottom, "", pyramid)
}

func dfsT(bottom, above string, pyramid map[string][]string) bool {
	if len(bottom) == 2 && len(above) == 1 {
		return true
	}
	if len(bottom) == len(above)+1 {
		return dfsT(above, "", pyramid)
	}
	base := bottom[len(above) : len(above)+2]
	if data, ok := pyramid[base]; ok {
		for _, key := range data {
			if dfsT(bottom, above+key, pyramid) {
				return true
			}
		}
	}
	return false
}
