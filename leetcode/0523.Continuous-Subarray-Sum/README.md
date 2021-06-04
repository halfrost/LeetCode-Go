# [523. Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/)


## 题目

Given an integer array `nums` and an integer `k`, return `true` *if* `nums` *has a continuous subarray of size **at least two** whose elements sum up to a multiple of* `k`*, or* `false` *otherwise*.

An integer `x` is a multiple of `k` if there exists an integer `n` such that `x = n * k`. `0` is **always** a multiple of `k`.

**Example 1:**

```
Input: nums = [23,2,4,6,7], k = 6
Output: true
Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.
```

**Example 2:**

```
Input: nums = [23,2,6,4,7], k = 6
Output: true
Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.
```

**Example 3:**

```
Input: nums = [23,2,6,4,7], k = 13
Output: false
```

**Constraints:**

- `1 <= nums.length <= 105`
- `0 <= nums[i] <= 109`
- `0 <= sum(nums[i]) <= 231 - 1`
- `1 <= k <= 231 - 1`

## 题目大意

给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

- 子数组大小至少为 2 ，且
- 子数组元素总和为 k 的倍数。

如果存在，返回 true ；否则，返回 false 。如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。

## 解题思路

- 简单题。题目只要求是否存在，不要求找出所有解。用一个变量 sum 记录累加和。子数组的元素和可以用前缀和相减得到，例如 [i,j] 区间内的元素和，可以由 prefixSum[j] - prefixSum[i] 得到。当 prefixSums[j]−prefixSums[i] 为 k 的倍数时，prefixSums[i] 和 prefixSums[j] 除以 k 的余数相同。因此只需要计算每个下标对应的前缀和除以 k 的余数即可，使用 map 存储每个余数第一次出现的下标即可。在 map 中如果存在相同余数的 key，代表当前下标和 map 中这个 key 记录的下标可以满足总和为 k 的倍数这一条件。再判断一下能否满足大小至少为 2 的条件即可。用 2 个下标相减，长度大于等于 2 即满足条件，可以输出 true。

## 代码

```go
package leetcode

func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i, n := range nums {
		sum += n
		if r, ok := m[sum%k]; ok {
			if i-2 >= r {
				return true
			}
		} else {
			m[sum%k] = i
		}
	}
	return false
}
```