# [1658. Minimum Operations to Reduce X to Zero](https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/)


## 题目

You are given an integer array `nums` and an integer `x`. In one operation, you can either remove the leftmost or the rightmost element from the array `nums` and subtract its value from `x`. Note that this **modifies** the array for future operations.

Return *the **minimum number** of operations to reduce* `x` *to **exactly*** `0` *if it's possible, otherwise, return* `1`.

**Example 1:**

```
Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.

```

**Example 2:**

```
Input: nums = [5,6,7,8,9], x = 4
Output: -1

```

**Example 3:**

```
Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.

```

**Constraints:**

- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 104`
- `1 <= x <= 109`

## 题目大意

给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。

## 解题思路

- 给定一个数组 nums 和一个整数 x，要求从数组两端分别移除一些数，使得这些数加起来正好等于整数 x，要求输出最小操作数。
- 要求输出最小操作数，即数组两头的数字个数最少，并且加起来和正好等于整数 x。由于在数组的两头，用 2 个指针分别操作不太方便。我当时解题的时候的思路是把它变成循环数组，这样两边的指针就在一个区间内了。利用滑动窗口找到一个最小的窗口，使得窗口内的累加和等于整数 k。这个方法可行，但是代码挺多的。
- 有没有更优美的方法呢？有的。要想两头的长度最少，也就是中间这段的长度最大。这样就转换成直接在数组上使用滑动窗口求解，累加和等于一个固定值的连续最长的子数组。
- 和这道题类似思路的题目，209，1040(循环数组)，325。强烈推荐这 3 题。

## 代码

```go
package leetcode

func minOperations(nums []int, x int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	target := total - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	left, right, sum, res := 0, 0, 0, -1
	for right < len(nums) {
		if sum < target {
			sum += nums[right]
			right++
		}
		for sum >= target {
			if sum == target {
				res = max(res, right-left)
			}
			sum -= nums[left]
			left++
		}
	}
	if res == -1 {
		return -1
	}
	return len(nums) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```