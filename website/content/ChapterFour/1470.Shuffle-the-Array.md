# [1470. Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/)

## 题目

Given the array `nums` consisting of `2n` elements in the form `[x1,x2,...,xn,y1,y2,...,yn]`.

*Return the array in the form* `[x1,y1,x2,y2,...,xn,yn]`.

**Example 1**:

```
Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7] 
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].

```

**Example 2**:

```
Input: nums = [1,2,3,4,4,3,2,1], n = 4
Output: [1,4,2,3,3,2,4,1]

```

**Example 3**:

```
Input: nums = [1,1,2,2], n = 2
Output: [1,2,1,2]

```

**Constraints**:

- `1 <= n <= 500`
- `nums.length == 2n`
- `1 <= nums[i] <= 10^3`

## 题目大意

给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。

## 解题思路

- 给定一个 2n 的数组，把后 n 个元素插空放到前 n 个元素里面。输出最终完成的数组。
- 简单题，按照题意插空即可。

## 代码

```go

package leetcode

func shuffle(nums []int, n int) []int {
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result
}

```