package leetcode

func halvesAreAlike(s string) bool {
	return numVowels(s[len(s)/2:]) == numVowels(s[:len(s)/2])
}

func numVowels(x string) int {
	res := 0
	for _, c := range x {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			res++
		}
	}
	return res
}
