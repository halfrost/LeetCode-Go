package leetcode

// 解法一
func partition131(s string) [][]string {
	if s == "" {
		return [][]string{}
	}
	res, pal := [][]string{}, []string{}
	findPalindrome(s, 0, "", true, pal, &res)
	return res
}

func findPalindrome(str string, index int, s string, isPal bool, pal []string, res *[][]string) {
	if index == len(str) {
		if isPal {
			tmp := make([]string, len(pal))
			copy(tmp, pal)
			*res = append(*res, tmp)
		}
		return
	}
	if index == 0 {
		s = string(str[index])
		pal = append(pal, s)
		findPalindrome(str, index+1, s, isPal && isPalindrome131(s), pal, res)
	} else {
		temp := pal[len(pal)-1]
		s = pal[len(pal)-1] + string(str[index])
		pal[len(pal)-1] = s
		findPalindrome(str, index+1, s, isPalindrome131(s), pal, res)
		pal[len(pal)-1] = temp
		if isPalindrome131(temp) {
			pal = append(pal, string(str[index]))
			findPalindrome(str, index+1, temp, isPal && isPalindrome131(temp), pal, res)
			pal = pal[:len(pal)-1]

		}
	}
	return
}

func isPalindrome131(s string) bool {
	slen := len(s)
	for i, j := 0, slen-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 解法二
func partition131_1(s string) [][]string {
	result := [][]string{}
	size := len(s)
	if size == 0 {
		return result
	}
	current := make([]string, 0, size)
	dfs131(s, 0, current, &result)
	return result
}

func dfs131(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)
	if start == end {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}
	for i := start; i < end; i++ {
		if isPal(s, start, i) {
			dfs131(s, i+1, append(cur, s[start:i+1]), result)
		}
	}
}

func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
