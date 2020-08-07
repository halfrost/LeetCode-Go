package leetcode

import "sort"

// 解法一 桶排序，时间复杂度 O(n^2)
func relativeSortArray(A, B []int) []int {
	count := [1001]int{}
	for _, a := range A {
		count[a]++
	}
	res := make([]int, 0, len(A))
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}
	return res
}

// 解法二 模拟，时间复杂度 O(n^2)
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	leftover, m, res := []int{}, make(map[int]int), []int{}
	for _, v := range arr1 {
		m[v]++
	}
	for _, s := range arr2 {
		count := m[s]
		for i := 0; i < count; i++ {
			res = append(res, s)
		}
		m[s] = 0
	}
	for v, count := range m {
		for i := 0; i < count; i++ {
			leftover = append(leftover, v)
		}
	}
	sort.Ints(leftover)
	res = append(res, leftover...)
	return res
}
