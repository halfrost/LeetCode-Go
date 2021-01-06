package leetcode

func findKthPositive(arr []int, k int) int {
	positive, index := 1, 0
	for index < len(arr) {
		if arr[index] != positive {
			k--
		} else {
			index++
		}
		if k == 0 {
			break
		}
		positive++
	}
	if k != 0 {
		positive += k - 1
	}
	return positive
}
