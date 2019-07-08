package leetcode

// 解法一
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	charMap, mp, result := map[uint8]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}, make(map[uint32]int, 0), []string{}
	var cur uint32
	for i := 0; i < 9; i++ { // 前9位，忽略
		cur = cur<<2 | charMap[s[i]]
	}
	for i := 9; i < len(s); i++ {
		cur = ((cur << 2) & 0xFFFFF) | charMap[s[i]]
		if mp[cur] == 0 {
			mp[cur] = 1
		} else if mp[cur] == 1 { // >2，重复
			mp[cur] = 2
			result = append(result, s[i-9:i+1])
		}
	}
	return result
}

// 解法二
func findRepeatedDnaSequences1(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	ans, cache := make([]string, 0), make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		curr := string(s[i : i+10])
		if cache[curr] == 1 {
			ans = append(ans, curr)
		}
		cache[curr]++
	}
	return ans
}
