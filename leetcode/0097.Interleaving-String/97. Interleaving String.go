package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	visited := make(map[int]bool)
	return dfs(s1, s2, s3, 0, 0, visited)
}

func dfs(s1, s2, s3 string, p1, p2 int, visited map[int]bool) bool {
	if p1+p2 == len(s3) {
		return true
	}
	if _, ok := visited[(p1*len(s3))+p2]; ok {
		return false
	}
	visited[(p1*len(s3))+p2] = true
	var match1, match2 bool
	if p1 < len(s1) && s3[p1+p2] == s1[p1] {
		match1 = true
	}
	if p2 < len(s2) && s3[p1+p2] == s2[p2] {
		match2 = true
	}
	if match1 && match2 {
		return dfs(s1, s2, s3, p1+1, p2, visited) || dfs(s1, s2, s3, p1, p2+1, visited)
	} else if match1 {
		return dfs(s1, s2, s3, p1+1, p2, visited)
	} else if match2 {
		return dfs(s1, s2, s3, p1, p2+1, visited)
	} else {
		return false
	}
}
