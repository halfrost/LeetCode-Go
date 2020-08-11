package leetcode

func findShortestSubArray(nums []int) int {
	frequency, maxFreq, smallest := map[int][]int{}, 0, len(nums)
	for i, num := range nums {
		if _, found := frequency[num]; !found {
			frequency[num] = []int{1, i, i}
		} else {
			frequency[num][0]++
			frequency[num][2] = i
		}
		if maxFreq < frequency[num][0] {
			maxFreq = frequency[num][0]
		}
	}
	for _, indices := range frequency {
		if indices[0] == maxFreq {
			if smallest > indices[2]-indices[1]+1 {
				smallest = indices[2] - indices[1] + 1
			}
		}
	}
	return smallest
}
