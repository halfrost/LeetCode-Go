package leetcode

import "math"

func commonChars(A []string) []string {
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = math.MaxUint16
	}
	cntInWord := [26]int{}
	for _, word := range A {
		for _, char := range []byte(word) { // compiler trick - here we will not allocate new memory
			cntInWord[char-'a']++
		}
		for i := 0; i < 26; i++ {
			// 缩小频次，使得统计的公共频次更加准确
			if cntInWord[i] < cnt[i] {
				cnt[i] = cntInWord[i]
			}
		}
		// 重置状态
		for i := range cntInWord {
			cntInWord[i] = 0
		}
	}
	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			result = append(result, string(i+'a'))
		}
	}
	return result
}
