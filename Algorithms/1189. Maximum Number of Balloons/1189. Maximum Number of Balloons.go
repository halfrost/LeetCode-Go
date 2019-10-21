package leetcode

func maxNumberOfBalloons(text string) int {
	fre := make([]int, 26)
	for _, t := range text {
		fre[t-'a']++
	}
	// 字符 b 的频次是数组下标 1 对应的元素值
	// 字符 a 的频次是数组下标 0 对应的元素值
	// 字符 l 的频次是数组下标 11 对应的元素值，这里有 2 个 l，所以元素值需要除以 2
	// 字符 o 的频次是数组下标 14 对应的元素值，这里有 2 个 o，所以元素值需要除以 2
	// 字符 n 的频次是数组下标 13 对应的元素值
	return min(fre[1], min(fre[0], min(fre[11]/2, min(fre[14]/2, fre[13]))))
}
