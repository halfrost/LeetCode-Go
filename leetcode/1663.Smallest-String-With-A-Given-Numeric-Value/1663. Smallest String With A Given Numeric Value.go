package leetcode

// 解法一 贪心
func getSmallestString(n int, k int) string {
	str, i, j := make([]byte, n), 0, 0
	for i = n - 1; i <= k-26; i, k = i-1, k-26 {
		str[i] = 'z'
	}
	if i >= 0 {
		str[i] = byte('a' + k - 1 - i)
		for ; j < i; j++ {
			str[j] = 'a'
		}
	}
	return string(str)
}

// 解法二 DFS
func getSmallestString1(n int, k int) string {
	if n == 0 {
		return ""
	}
	res, c := "", []byte{}
	findSmallestString(0, n, k, 0, c, &res)
	return res
}

func findSmallestString(value int, length, k, index int, str []byte, res *string) {
	if len(str) == length && value == k {
		tmp := string(str)
		if (*res) == "" {
			*res = tmp
		}
		if tmp < *res && *res != "" {
			*res = tmp
		}
		return
	}
	if len(str) >= index && (*res) != "" && str[index-1] > (*res)[index-1] {
		return
	}
	for j := 0; j < 26; j++ {
		if k-value > (length-len(str))*26 || value > k {
			return
		}
		str = append(str, byte(int('a')+j))
		value += j + 1
		findSmallestString(value, length, k, index+1, str, res)
		str = str[:len(str)-1]
		value -= j + 1

	}
}
