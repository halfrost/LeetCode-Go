package leetcode

func replaceElements(arr []int) []int {
	j, temp := -1, 0
	for i := len(arr) - 1; i >= 0; i-- {
		temp = arr[i]
		arr[i] = j
		j = max(j, temp)
	}
	return arr
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
