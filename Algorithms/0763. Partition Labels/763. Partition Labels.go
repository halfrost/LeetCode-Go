package leetcode

// 解法一
func partitionLabels(S string) []int {
	var lastIndexOf [26]int
	for i, v := range S {
		lastIndexOf[v-'a'] = i
	}

	var arr []int
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = lastIndexOf[S[start]-'a']
		for i := start; i < end; i++ {
			if end < lastIndexOf[S[i]-'a'] {
				end = lastIndexOf[S[i]-'a']
			}
		}
		arr = append(arr, end-start+1)
	}
	return arr
}

// 解法二
func partitionLabels1(S string) []int {
	visit, counter, res, sum, lastLength := make([]int, 26), map[byte]int{}, []int{}, 0, 0
	for i := 0; i < len(S); i++ {
		counter[S[i]]++
	}

	for i := 0; i < len(S); i++ {
		counter[S[i]]--
		visit[S[i]-'a'] = 1
		sum = 0
		for j := 0; j < 26; j++ {
			if visit[j] == 1 {
				sum += counter[byte('a'+j)]
			}
		}
		if sum == 0 {
			res = append(res, i+1-lastLength)
			lastLength += i + 1 - lastLength
		}
	}
	return res
}
