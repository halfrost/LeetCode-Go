package leetcode

// 解法一 普通做法
func dailyTemperatures(T []int) []int {
	res, j := make([]int, len(T)), 0
	for i := 0; i < len(T); i++ {
		for j = i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

// 解法二 单调栈
func dailyTemperatures1(T []int) []int {
	res := make([]int, len(T))
	var toCheck []int
	for i, t := range T {
		for len(toCheck) > 0 && T[toCheck[len(toCheck)-1]] < t {
			idx := toCheck[len(toCheck)-1]
			res[idx] = i - idx
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, i)
	}
	return res
}
