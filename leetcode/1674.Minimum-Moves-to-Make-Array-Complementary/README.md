# [1674. Minimum Moves to Make Array Complementary](https://leetcode.com/problems/minimum-moves-to-make-array-complementary/)

## 题目

You are given an integer array `nums` of **even** length `n` and an integer `limit`. In one move, you can replace any integer from `nums` with another integer between `1` and `limit`, inclusive.

The array `nums` is **complementary** if for all indices `i` (**0-indexed**), `nums[i] + nums[n - 1 - i]` equals the same number. For example, the array `[1,2,3,4]` is complementary because for all indices `i`, `nums[i] + nums[n - 1 - i] = 5`.

Return the ***minimum** number of moves required to make* `nums` ***complementary***.

**Example 1:**

```
Input: nums = [1,2,4,3], limit = 4
Output: 1
Explanation: In 1 move, you can change nums to [1,2,2,3] (underlined elements are changed).
nums[0] + nums[3] = 1 + 3 = 4.
nums[1] + nums[2] = 2 + 2 = 4.
nums[2] + nums[1] = 2 + 2 = 4.
nums[3] + nums[0] = 3 + 1 = 4.
Therefore, nums[i] + nums[n-1-i] = 4 for every i, so nums is complementary.
```

**Example 2:**

```
Input: nums = [1,2,2,1], limit = 2
Output: 2
Explanation: In 2 moves, you can change nums to [2,2,2,2]. You cannot change any number to 3 since 3 > limit.
```

**Example 3:**

```
Input: nums = [1,2,1,2], limit = 2
Output: 0
Explanation: nums is already complementary.
```

**Constraints:**

- `n == nums.length`
- `2 <= n <= 105`
- `1 <= nums[i] <= limit <= 105`
- `n` is even.

## 题目大意

给你一个长度为 偶数 n 的整数数组 nums 和一个整数 limit 。每一次操作，你可以将 nums 中的任何整数替换为 1 到 limit 之间的另一个整数。

如果对于所有下标 i（下标从 0 开始），nums[i] + nums[n - 1 - i] 都等于同一个数，则数组 nums 是 互补的 。例如，数组 [1,2,3,4] 是互补的，因为对于所有下标 i ，nums[i] + nums[n - 1 - i] = 5 。

返回使数组 互补 的 最少 操作次数。

## 解题思路

- 这一题考察的是差分数组。通过分析题意，可以得出，针对每一个 `sum` 的取值范围是 `[2, 2* limt]`，定义 `a = min(nums[i], nums[n - i - 1])`，`b = max(nums[i], nums[n - i - 1])`，在这个区间内，又可以细分成 5 个区间，`[2, a + 1)`，`[a + 1, a + b)`，`[a + b + 1, a + b + 1)`，`[a + b + 1, b + limit + 1)`，`[b + limit + 1, 2 * limit)`，在这 5 个区间内使得数组互补的最小操作次数分别是 `2(减少 a, 减少 b)`，`1(减少 b)`，`0(不用操作)`，`1(增大 a)`，`+2(增大 a, 增大 b)`，换个表达方式，按照扫描线从左往右扫描，在这 5 个区间内使得数组互补的最小操作次数叠加变化分别是 `+2(减少 a, 减少 b)`，`-1(减少 a)`，`-1(不用操作)`，`+1(增大 a)`，`+1(增大 a, 增大 b)`，利用这前后两个区间的关系，就可以构造一个差分数组。差分数组反应的是前后两者的关系。如果想求得 0 ~ n 的总关系，只需要求一次前缀和即可。
- 这道题要求输出最少的操作次数，所以利用差分数组 + 前缀和，累加前缀和的同时维护最小值。从左往右扫描完一遍以后，输出最小值即可。

## 代码

```go
package leetcode

func minMoves(nums []int, limit int) int {
	diff := make([]int, limit*2+2) // nums[i] <= limit, b+limit+1 is maximum limit+limit+1
	for j := 0; j < len(nums)/2; j++ {
		a, b := min(nums[j], nums[len(nums)-j-1]), max(nums[j], nums[len(nums)-j-1])
		// using prefix sum: most interesting point, and is the key to reduce complexity
		diff[2] += 2
		diff[a+1]--
		diff[a+b]--
		diff[a+b+1]++
		diff[b+limit+1]++
	}
	cur, res := 0, len(nums)
	for i := 2; i <= 2*limit; i++ {
		cur += diff[i]
		res = min(res, cur)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```