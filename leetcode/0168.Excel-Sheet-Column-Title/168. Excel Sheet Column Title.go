package leetcode

func convertToTitle(n int) string {
	result := []byte{}
	for n > 0 {
		result = append(result, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
