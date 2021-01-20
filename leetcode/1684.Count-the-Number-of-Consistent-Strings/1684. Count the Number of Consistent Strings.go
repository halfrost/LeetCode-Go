package leetcode

func countConsistentStrings(allowed string, words []string) int {
	allowedMap, res, flag := map[rune]int{}, 0, true
	for _, str := range allowed {
		allowedMap[str]++
	}
	for i := 0; i < len(words); i++ {
		flag = true
		for j := 0; j < len(words[i]); j++ {
			if _, ok := allowedMap[rune(words[i][j])]; !ok {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}
