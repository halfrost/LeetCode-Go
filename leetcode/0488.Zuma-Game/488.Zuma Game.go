package leetcode

func findMinStep(board string, hand string) int {
	q := [][]string{{board, hand}}
	mp := make(map[string]bool)
	minStep := 0
	for len(q) > 0 {
		length := len(q)
		minStep++
		for length > 0 {
			length--
			cur := q[0]
			q = q[1:]
			curB, curH := cur[0], cur[1]
			for i := 0; i < len(curB); i++ {
				for j := 0; j < len(curH); j++ {
					curB2 := del3(curB[0:i] + string(curH[j]) + curB[i:])
					curH2 := curH[0:j] + curH[j+1:]
					if len(curB2) == 0 {
						return minStep
					}
					if _, ok := mp[curB2+curH2]; ok {
						continue
					}
					mp[curB2+curH2] = true
					q = append(q, []string{curB2, curH2})
				}
			}
		}
	}
	return -1
}

func del3(str string) string {
	cnt := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			cnt++
		} else {
			if cnt >= 3 {
				return del3(str[0:i-cnt] + str[i:])
			}
			cnt = 1
		}
	}
	if cnt >= 3 {
		return str[0 : len(str)-cnt]
	}
	return str
}
