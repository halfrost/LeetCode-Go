package leetcode

import (
	"math/rand"
	"sort"
)

// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 解法二 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
// 时间复杂度 O(n)，空间复杂度 O(log n)，最坏时间复杂度为 O(n^2)，空间复杂度 O(n)
func findKthLargest(nums []int, k int) int {
	m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
	return selectSmallest(nums, 0, len(nums)-1, m)
}

func selectSmallest(nums []int, l, r, i int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums[q]
	}
	if i < k {
		return selectSmallest(nums, l, q-1, i)
	} else {
		return selectSmallest(nums, q+1, r, i-k)
	}
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1) // 此处为优化，使得时间复杂度期望降为 O(n)，最坏时间复杂度为 O(n^2)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	// nums[l..i] <= nums[r]
	// nums[i+1..j-1] > nums[r]
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// 扩展题 剑指 Offer 40. 最小的 k 个数
func getLeastNumbers(arr []int, k int) []int {
	return selectSmallest1(arr, 0, len(arr)-1, k)[:k]
}

// 和 selectSmallest 实现完全一致，只是返回值不用再截取了，直接返回 nums 即可
func selectSmallest1(nums []int, l, r, i int) []int {
	if l >= r {
		return nums
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums
	}
	if i < k {
		return selectSmallest1(nums, l, q-1, i)
	} else {
		return selectSmallest1(nums, q+1, r, i-k)
	}
}
