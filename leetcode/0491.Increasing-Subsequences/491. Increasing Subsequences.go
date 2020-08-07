package leetcode

func findSubsequences(nums []int) [][]int {
	c, visited, res := []int{}, map[int]bool{}, [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		} else {
			visited[nums[i]] = true
			generateIncSubsets(nums, i, c, &res)
		}
	}
	return res
}

func generateIncSubsets(nums []int, current int, c []int, res *[][]int) {
	c = append(c, nums[current])
	if len(c) >= 2 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
	}
	visited := map[int]bool{}
	for i := current + 1; i < len(nums); i++ {
		if nums[current] <= nums[i] {
			if _, ok := visited[nums[i]]; ok {
				continue
			} else {
				visited[nums[i]] = true
				generateIncSubsets(nums, i, c, res)
			}
		}
	}
	c = c[:len(c)-1]
	return
}
