package leetcode

// 解法一 这一题可以利用数组有序的特性
func twoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

// 解法二 不管数组是否有序，空间复杂度比上一种解法要多 O(n)
func twoSum167_1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if _, ok := m[another]; ok {
			return []int{m[another] + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}
