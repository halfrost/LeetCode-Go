package leetcode

func checkRecord(s string) bool {
	numsA := 0
	maxL := 0
	numsL := 0
	for _, v := range s {
		if v == 'L' {
			numsL++
		} else {
			if numsL > maxL {
				maxL = numsL
			}
			numsL = 0
			if v == 'A' {
				numsA++
			}
		}
	}
	if numsL > maxL {
		maxL = numsL
	}
	return numsA < 2 && maxL < 3
}
