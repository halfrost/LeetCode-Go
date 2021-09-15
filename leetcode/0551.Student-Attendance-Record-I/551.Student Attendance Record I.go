package leetcode

func checkRecord(s string) bool {
	numsA, maxL, numsL := 0, 0, 0
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
