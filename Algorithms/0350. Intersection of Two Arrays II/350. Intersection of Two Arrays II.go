package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}
