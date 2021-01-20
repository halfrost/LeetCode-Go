# [1685. Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/)


## 题目

You are given an integer array `nums` sorted in **non-decreasing** order.

Build and return *an integer array* `result` *with the same length as* `nums` *such that* `result[i]` *is equal to the **summation of absolute differences** between* `nums[i]` *and all the other elements in the array.*

In other words, `result[i]` is equal to `sum(|nums[i]-nums[j]|)` where `0 <= j < nums.length` and `j != i` (**0-indexed**).

**Example 1:**

```
Input: nums = [2,3,5]
Output: [4,3,5]
Explanation: Assuming the arrays are 0-indexed, then
result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4,
result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3,
result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5.
```

**Example 2:**

```
Input: nums = [1,4,6,8,10]
Output: [24,15,13,15,21]
```

**Constraints:**

- `2 <= nums.length <= 105`
- `1 <= nums[i] <= nums[i + 1] <= 104`

## 题目大意

给你一个 非递减 有序整数数组 `nums` 。请你建立并返回一个整数数组 `result`，它跟 `nums` 长度相同，且`result[i]` 等于 `nums[i]` 与数组中所有其他元素差的绝对值之和。换句话说， `result[i]` 等于 `sum(|nums[i]-nums[j]|)` ，其中 `0 <= j < nums.length` 且 `j != i` （下标从 0 开始）。

## 解题思路

- 利用前缀和思路解题。题目中说明了是有序数组，所以在计算绝对值的时候可以拆开绝对值符号。假设要计算当前 `result[i]`，以 `i` 为界，把原数组 `nums` 分成了 3 段。`nums[0 ~ i-1]` 和 `nums[i+1 ~ n]`，前面一段 `nums[0 ~ i-1]` 中的每个元素都比 `nums[i]` 小，拆掉绝对值以后，`sum(|nums[i]-nums[j]|) = nums[i] * i - prefixSum[0 ~ i-1]`，后面一段 `nums[i+1 ~ n]` 中的每个元素都比 `nums[i]` 大，拆掉绝对值以后，`sum(|nums[i]-nums[j]|) = prefixSum[i+1 ~ n] - nums[i] * (n - 1 - i)`。特殊的情况，`i = 0` 和 `i = n` 的情况特殊处理一下就行。

## 代码

```go
package leetcode

//解法一 优化版 prefixSum + sufixSum
func getSumAbsoluteDifferences(nums []int) []int {
	size := len(nums)
	sufixSum := make([]int, size)
	sufixSum[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		sufixSum[i] = sufixSum[i+1] + nums[i]
	}
	ans, preSum := make([]int, size), 0
	for i := 0; i < size; i++ {
		// 后面可以加到的值
		res, sum := 0, sufixSum[i]-nums[i]
		res += (sum - (size-i-1)*nums[i])
		// 前面可以加到的值
		res += (i*nums[i] - preSum)
		ans[i] = res
		preSum += nums[i]
	}
	return ans
}

// 解法二 prefixSum
func getSumAbsoluteDifferences1(nums []int) []int {
	preSum, res, sum := []int{}, []int{}, nums[0]
	preSum = append(preSum, nums[0])
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		preSum = append(preSum, sum)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res = append(res, preSum[len(nums)-1]-preSum[0]-nums[i]*(len(nums)-1))
		} else if i > 0 && i < len(nums)-1 {
			res = append(res, preSum[len(nums)-1]-preSum[i]-preSum[i-1]+nums[i]*i-nums[i]*(len(nums)-1-i))
		} else {
			res = append(res, nums[i]*len(nums)-preSum[len(nums)-1])
		}
	}
	return res
}
```