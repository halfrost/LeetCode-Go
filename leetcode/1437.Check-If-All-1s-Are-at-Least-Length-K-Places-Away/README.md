# [1437. Check If All 1's Are at Least Length K Places Away](https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/)


## 题目

Given an array `nums` of 0s and 1s and an integer `k`, return `True` if all 1's are at least `k` places away from each other, otherwise return `False`.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/04/15/sample_1_1791.png](https://assets.leetcode.com/uploads/2020/04/15/sample_1_1791.png)

```
Input: nums = [1,0,0,0,1,0,0,1], k = 2
Output: true
Explanation: Each of the 1s are at least 2 places away from each other.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/04/15/sample_2_1791.png](https://assets.leetcode.com/uploads/2020/04/15/sample_2_1791.png)

```
Input: nums = [1,0,0,1,0,1], k = 2
Output: false
Explanation: The second 1 and third 1 are only one apart from each other.
```

**Example 3:**

```
Input: nums = [1,1,1,1,1], k = 0
Output: true
```

**Example 4:**

```
Input: nums = [0,1,0,1], k = 1
Output: true
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `0 <= k <= nums.length`
- `nums[i]` is `0` or `1`

## 题目大意

给你一个由若干 0 和 1 组成的数组 nums 以及整数 k。如果所有 1 都至少相隔 k 个元素，则返回 True ；否则，返回 False 。

## 解题思路

- 简单题。扫描一遍数组，遇到 1 的时候比较前一个 1 的下标索引，如果相隔小于 k 则返回 false。如果大于等于 k 就更新下标索引，继续循环。循环结束输出 true 即可。

## 代码

```go
package leetcode

func kLengthApart(nums []int, k int) bool {
	prevIndex := -1
	for i, num := range nums {
		if num == 1 {
			if prevIndex != -1 && i-prevIndex-1 < k {
				return false
			}
			prevIndex = i
		}
	}
	return true
}
```