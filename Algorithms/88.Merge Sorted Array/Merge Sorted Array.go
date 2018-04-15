package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	// 这里不需要，因为测试数据考虑到了第一个数组的空间问题
	// for index := 0; index < n; index++ {
	// 	nums1 = append(nums1, nums2[index])
	// }
	i := m - 1
	j := n - 1
	k := m + n - 1
	// 从后面往前放，只需要循环一次即可
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
