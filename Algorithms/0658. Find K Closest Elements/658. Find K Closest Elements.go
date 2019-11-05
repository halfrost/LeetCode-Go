package leetcode

import "sort"

// 解法一 库函数二分搜索
func findClosestElements(arr []int, k int, x int) []int {
	return arr[sort.Search(len(arr)-k, func(i int) bool { return x-arr[i] <= arr[i+k]-x }):][:k]
}

// 解法二 手撸二分搜索
func findClosestElements1(arr []int, k int, x int) []int {
	low, high := 0, len(arr)-k
	for low < high {
		mid := low + (high-low)>>1
		if x-arr[mid] > arr[mid+k]-x {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return arr[low : low+k]
}
