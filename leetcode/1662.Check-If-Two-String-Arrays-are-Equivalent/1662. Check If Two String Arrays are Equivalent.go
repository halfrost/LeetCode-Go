package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1, str2 := "", ""
	for i := 0; i < len(word1); i++ {
		str1 += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		str2 += word2[i]
	}
	return str1 == str2
}
