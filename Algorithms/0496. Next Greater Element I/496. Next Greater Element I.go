package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res, reocrd, flag := []int{}, map[int]int{}, false
	for i, v := range nums2 {
		reocrd[v] = i
	}
	for i := 0; i < len(nums1); i++ {
		flag = false
		for j := reocrd[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				flag = true
				break
			}
		}
		if flag == false {
			res = append(res, -1)
		}
	}
	return res
}
