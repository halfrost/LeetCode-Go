package leetcode

func maxProduct318(words []string) int {
	if words == nil || len(words) == 0 {
		return 0
	}
	length, value, maxProduct := len(words), make([]int, len(words)), 0
	for i := 0; i < length; i++ {
		tmp := words[i]
		value[i] = 0
		for j := 0; j < len(tmp); j++ {
			value[i] |= 1 << (tmp[j] - 'a')
		}
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (value[i]&value[j]) == 0 && (len(words[i])*len(words[j]) > maxProduct) {
				maxProduct = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProduct
}
