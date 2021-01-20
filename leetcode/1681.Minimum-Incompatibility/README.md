# [1681. Minimum Incompatibility](https://leetcode.com/problems/minimum-incompatibility/)


## 题目

You are given an integer array `nums` and an integer `k`. You are asked to distribute this array into `k` subsets of **equal size** such that there are no two equal elements in the same subset.

A subset's **incompatibility** is the difference between the maximum and minimum elements in that array.

Return *the **minimum possible sum of incompatibilities** of the* `k` *subsets after distributing the array optimally, or return* `-1` *if it is not possible.*

A subset is a group integers that appear in the array with no particular order.

**Example 1:**

```
Input: nums = [1,2,1,4], k = 2
Output: 4
Explanation: The optimal distribution of subsets is [1,2] and [1,4].
The incompatibility is (2-1) + (4-1) = 4.
Note that [1,1] and [2,4] would result in a smaller sum, but the first subset contains 2 equal elements.
```

**Example 2:**

```
Input: nums = [6,3,8,1,3,1,2,2], k = 4
Output: 6
Explanation: The optimal distribution of subsets is [1,2], [2,3], [6,8], and [1,3].
The incompatibility is (2-1) + (3-2) + (8-6) + (3-1) = 6.

```

**Example 3:**

```
Input: nums = [5,3,3,6,3,3], k = 3
Output: -1
Explanation: It is impossible to distribute nums into 3 subsets where no two elements are equal in the same subset.

```

**Constraints:**

- `1 <= k <= nums.length <= 16`
- `nums.length` is divisible by `k`
- `1 <= nums[i] <= nums.length`

## 题目大意

给你一个整数数组 nums 和一个整数 k 。你需要将这个数组划分到 k 个相同大小的子集中，使得同一个子集里面没有两个相同的元素。一个子集的 不兼容性 是该子集里面最大值和最小值的差。

请你返回将数组分成 k 个子集后，各子集 **不兼容性** 的 **和** 的 **最小值** ，如果无法分成分成 k 个子集，返回 -1 。子集的定义是数组中一些数字的集合，对数字顺序没有要求。

## 解题思路

- 读完题最直白的思路就是 DFS。做法类似第 77 题。这里就不赘述了。可以见第 77 题题解。
- 这一题还需要用到贪心的思想。每次取数都取最小的数。这样可以不会让最大数和最小数在一个集合中。由于每次取数都是取最小的，那么能保证不兼容性每次都尽量最小。于是在 order 数组中定义取数的顺序。然后再把数组从小到大排列。这样每次按照 order 顺序取数，都是取的最小值。
- 正常的 DFS 写完提交，耗时是很长的。大概是 1532ms。如何优化到极致呢？这里需要加上 2 个剪枝条件。第一个剪枝条件比较简单，如果累计 sum 比之前存储的 res 大，那么直接 return，不需要继续递归了。第二个剪枝条件就非常重要了，可以一下子减少很多次递归。每次取数产生新的集合的时候，要从第一个最小数开始取，一旦取了，后面就不需要再循环递归了。举个例子，[1,2,3,4]，第一个数如果取 2，集合可以是 [[2,3],[1,4]] 或 [[2,4], [1,3]], 这个集合和[[1,3],[2,4]]、[[1,4], [2,3]] 情况一样。可以看到如果取出第一个最小值以后，后面的循环是不必要的了。所以在取下标为 0 的数的时候，递归到底层以后，返回就可以直接 break，不用接下去的循环了，接下去的循环和递归是不必要的。每组组内的顺序我们并不关心，只要最大值和最小值在分组内即可。另外组间顺序我们也不关心。所以可以把排列问题 O(n!) 时间复杂度降低到组合问题 O(2^n)。加了这 2 个剪枝条件以后，耗时就变成了 0ms 了。beats 100%

## 代码

```go
package leetcode

import (
	"math"
	"sort"
)

func minimumIncompatibility(nums []int, k int) int {
	sort.Ints(nums)
	eachSize, counts := len(nums)/k, make([]int, len(nums)+1)
	for i := range nums {
		counts[nums[i]]++
		if counts[nums[i]] > k {
			return -1
		}
	}
	orders := []int{}
	for i := range counts {
		orders = append(orders, i)
	}
	sort.Ints(orders)
	res := math.MaxInt32
	generatePermutation1681(nums, counts, orders, 0, 0, eachSize, &res, []int{})
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func generatePermutation1681(nums, counts, order []int, index, sum, eachSize int, res *int, current []int) {
	if len(current) > 0 && len(current)%eachSize == 0 {
		sum += current[len(current)-1] - current[len(current)-eachSize]
		index = 0
	}
	if sum >= *res {
		return
	}
	if len(current) == len(nums) {
		if sum < *res {
			*res = sum
		}
		return
	}
	for i := index; i < len(counts); i++ {
		if counts[order[i]] == 0 {
			continue
		}
		counts[order[i]]--
		current = append(current, order[i])
		generatePermutation1681(nums, counts, order, i+1, sum, eachSize, res, current)
		current = current[:len(current)-1]
		counts[order[i]]++
		// 这里是关键的剪枝
		if index == 0 {
			break
		}
	}
}
```