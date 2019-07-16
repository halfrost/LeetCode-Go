package leetcode

func findLHS(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, exist := res[num]; exist {
			res[num]++
			continue
		}
		res[num] = 1
	}
	longest := 0
	for k, c := range res {
		if n, exist := res[k+1]; exist {
			if c+n > longest {
				longest = c + n
			}
		}
	}
	return longest
}
