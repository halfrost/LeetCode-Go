package leetcode

import "strconv"

func getHint(secret string, guess string) string {
	cntA, cntB := 0, 0
	mpS := make(map[byte]int)
	var strG []byte
	n := len(secret)
	var ans string
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			cntA++
		} else {
			mpS[secret[i]] += 1
			strG = append(strG, guess[i])
		}
	}
	for _, v := range strG {
		if _, ok := mpS[v]; ok {
			if mpS[v] > 1 {
				mpS[v] -= 1
			} else {
				delete(mpS, v)
			}
			cntB++
		}
	}
	ans += strconv.Itoa(cntA) + "A" + strconv.Itoa(cntB) + "B"
	return ans
}
