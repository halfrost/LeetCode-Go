package leetcode

import (
	"sort"
)

// 解法一
func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	median := findKthLargest324(nums, (len(nums)+1)/2)
	n, i, left, right := len(nums), 0, 0, len(nums)-1

	for i <= right {
		if nums[indexMap(i, n)] > median {
			nums[indexMap(left, n)], nums[indexMap(i, n)] = nums[indexMap(i, n)], nums[indexMap(left, n)]
			left++
			i++
		} else if nums[indexMap(i, n)] < median {
			nums[indexMap(right, n)], nums[indexMap(i, n)] = nums[indexMap(i, n)], nums[indexMap(right, n)]
			right--
		} else {
			i++
		}
	}
}

func indexMap(index, n int) int {
	return (1 + 2*index) % (n | 1)
}

func findKthLargest324(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selection324(nums, 0, len(nums)-1, len(nums)-k)
}

func selection324(arr []int, l, r, k int) int {
	if l == r {
		return arr[l]
	}
	p := partition324(arr, l, r)

	if k == p {
		return arr[p]
	} else if k < p {
		return selection324(arr, l, p-1, k)
	} else {
		return selection324(arr, p+1, r, k)
	}
}

func partition324(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

// 解法二
func wiggleSort1(nums []int) {
	if len(nums) < 2 {
		return
	}
	array := make([]int, len(nums))
	copy(array, nums)
	sort.Ints(array)
	n := len(nums)
	left := (n+1)/2 - 1 // median index
	right := n - 1      // largest value index
	for i := 0; i < len(nums); i++ {
		// copy large values on odd indexes
		if i%2 == 1 {
			nums[i] = array[right]
			right--
		} else { // copy values decremeting from median on even indexes
			nums[i] = array[left]
			left--
		}
	}
}
