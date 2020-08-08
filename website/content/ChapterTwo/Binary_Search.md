---
title: Binary Search
type: docs
---

# Binary Search

- 二分搜索的经典写法。需要注意的三点：
	1. 循环退出条件，注意是 low <= high，而不是 low < high。
	2. mid 的取值，mid := low + (high-low)>>1
	3. low 和 high 的更新。low = mid + 1，high = mid - 1。

```go
func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
```

- 二分搜索的变种写法。有 4 个基本变种:
	1. 查找第一个与 target 相等的元素，时间复杂度 O(logn)
	2. 查找最后一个与 target 相等的元素，时间复杂度 O(logn)
	3. 查找第一个大于等于 target 的元素，时间复杂度 O(logn)
	4. 查找最后一个小于等于 target 的元素，时间复杂度 O(logn)

```go
// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

- 在基本有序的数组中用二分搜索。经典解法可以解，变种写法也可以写，常见的题型，在山峰数组中找山峰，在旋转有序数组中找分界点。第 33 题，第 81 题，第 153 题，第 154 题，第 162 题，第 852 题

```go
func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)>>1
		// 如果 mid 较大，则左侧存在峰值，high = m，如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if A[mid] > A[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
```

- max-min 最大值最小化问题。求在最小满足条件的情况下的最大值。第 410 题，第 875 题，第 1011 题，第 1283 题。

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|50. Pow(x, n) | [Go]({{< relref "/ChapterFour/0050.Powx-n.md" >}})| Medium | O(log n)| O(1)||
|69. Sqrt(x)  | [Go]({{< relref "/ChapterFour/0069.Sqrtx.md" >}})| Easy | O(log n)| O(1)||
|167. Two Sum II - Input array is sorted | [Go]({{< relref "/ChapterFour/0167.Two-Sum-II---Input-array-is-sorted.md" >}})| Easy | O(n)| O(1)||
|209. Minimum Size Subarray Sum  | [Go]({{< relref "/ChapterFour/0209.Minimum-Size-Subarray-Sum.md" >}})| Medium | O(n)| O(1)||
|222. Count Complete Tree Nodes  | [Go]({{< relref "/ChapterFour/0222.Count-Complete-Tree-Nodes.md" >}})| Medium | O(n)| O(1)||
|230. Kth Smallest Element in a BST  | [Go]({{< relref "/ChapterFour/0230.Kth-Smallest-Element-in-a-BST.md" >}})| Medium | O(n)| O(1)||
|287. Find the Duplicate Number   | [Go]({{< relref "/ChapterFour/0287.Find-the-Duplicate-Number.md" >}})| Easy | O(n)| O(1)|❤️|
|300. Longest Increasing Subsequence  | [Go]({{< relref "/ChapterFour/0300.Longest-Increasing-Subsequence.md" >}})| Medium | O(n log n)| O(n)||
|349. Intersection of Two Arrays  | [Go]({{< relref "/ChapterFour/0349.Intersection-of-Two-Arrays.md" >}})| Easy | O(n)| O(n) ||
|350. Intersection of Two Arrays II  | [Go]({{< relref "/ChapterFour/0350.Intersection-of-Two-Arrays-II.md" >}})| Easy | O(n)| O(n) ||
|392. Is Subsequence  | [Go]({{< relref "/ChapterFour/0392.Is-Subsequence.md" >}})| Medium | O(n)| O(1)||
|454. 4Sum II | [Go]({{< relref "/ChapterFour/0454.4Sum-II.md" >}})| Medium | O(n^2)| O(n) ||
|710. Random Pick with Blacklist | [Go]({{< relref "/ChapterFour/0710.Random-Pick-with-Blacklist.md" >}})| Hard | O(n)| O(n)  ||
|-----------------------------------------------|---------------------------------|--------------------------|-----------------------|-----------|--------|