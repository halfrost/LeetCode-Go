package leetcode

func countSegments(s string) int {
	segments := false
	cnt := 0
	for _, v := range s {
		if v == ' ' && segments {
			segments = false
			cnt += 1
		} else if v != ' ' {
			segments = true
		}
	}
	if segments {
		cnt++
	}
	return cnt
}
