# [376. Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/)


## 题目

Given an integer array `nums`, return *the length of the longest **wiggle sequence***.

A **wiggle sequence** is a sequence where the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

- For example, `[1, 7, 4, 9, 2, 5]` is a **wiggle sequence** because the differences `(6, -3, 5, -7, 3)` are alternately positive and negative.
- In contrast, `[1, 4, 7, 2, 5]` and `[1, 7, 4, 5, 5]` are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

A **subsequence** is obtained by deleting some elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

**Example 1:**

```
Input: nums = [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence.
```

**Example 2:**

```
Input: nums = [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
```

**Example 3:**

```
Input: nums = [1,2,3,4,5,6,7,8,9]
Output: 2
```

**Constraints:**

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`

**Follow up:** Could you solve this in `O(n)` time?

## 题目大意

如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。

## 解题思路

- 题目要求找到摆动序列最长的子序列。本题可以用贪心的思路，记录当前序列的上升和下降的趋势。扫描数组过程中，每扫描一个元素都判断是“峰”还是“谷”，根据前一个是“峰”还是“谷”做出对应的决定。利用贪心的思想找到最长的摆动子序列。

## 代码

```go
package leetcode

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	res := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		res = 2
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			res++
			prevDiff = diff
		}
	}
	return res
}
```