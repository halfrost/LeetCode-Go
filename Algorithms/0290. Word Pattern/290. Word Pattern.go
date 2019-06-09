package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	patternByte := []byte(pattern)
	if pattern == "" || len(patternByte) != len(strList) {
		return false
	}

	pMap := map[byte]string{}
	sMap := map[string]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}
