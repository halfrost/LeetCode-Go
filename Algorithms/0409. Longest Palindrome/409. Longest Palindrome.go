package leetcode

func longestPalindrome(s string) int {
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}
	answer := 0
	for _, v := range counter {
		answer += v / 2 * 2
		if answer%2 == 0 && v%2 == 1 {
			answer++
		}
	}
	return answer
}
