package leetcode

func isIsomorphic(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}

	pMap := map[byte]byte{}
	sMap := map[byte]byte{}
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
