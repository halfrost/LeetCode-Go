package leetcode

// 解法一 贪心
func getSmallestString(n int, k int) string {
	res := make([]rune, n)
	for i := n - 1; i >= 0; i-- {
		diff := k - i
		if diff >= 26 {
			// Need to add z
			res[i] = 'z'
			k = k - 26
		} else {
			res[i] = rune('a' + diff - 1)
			k = k - diff
		}
	}
	return string(res)
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

func getSmallestString2(n int, k int) string {
	var result = make([]byte, n)
	for pos := n - 1; pos >= 0; pos-- {
		add := min(k - pos, 26)
		result[pos] = byte(add) + 'a' - 1
		k -= add
	}
	return string(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
