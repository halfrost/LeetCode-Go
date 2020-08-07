package leetcode

func findPairs(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		return 0
	}
	var count int
	m := make(map[int]int, len(nums))
	for _, value := range nums {
		m[value]++
	}
	for key := range m {
		if k == 0 && m[key] > 1 {
			count++
			continue
		}
		if k > 0 && m[key+k] > 0 {
			count++
		}
	}
	return count
}
