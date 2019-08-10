package leetcode

func findLongestWord(s string, d []string) string {
	res := ""
	for i := 0; i < len(d); i++ {
		pointS := 0
		pointD := 0
		for pointS < len(s) && pointD < len(d[i]) {
			if s[pointS] == d[i][pointD] {
				pointD++
			}
			pointS++
		}
		if pointD == len(d[i]) && (len(res) < len(d[i]) || (len(res) == len(d[i]) && res > d[i])) {
			res = d[i]
		}
	}
	return res
}
