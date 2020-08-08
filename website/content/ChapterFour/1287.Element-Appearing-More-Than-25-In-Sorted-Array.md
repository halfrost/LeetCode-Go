# [1287. Element Appearing More Than 25% In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/)



## 题目

Given an integer array **sorted** in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time.

Return that integer.

**Example 1**:

```
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6
```

**Constraints**:

- `1 <= arr.length <= 10^4`
- `0 <= arr[i] <= 10^5`

## 题目大意

给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。请你找到并返回这个整数。

提示：

- 1 <= arr.length <= 10^4
- 0 <= arr[i] <= 10^5

## 解题思路

- 给出一个非递减的有序数组，要求输出出现次数超过数组元素总数 25% 的元素。
- 简单题，由于已经非递减有序了，所以只需要判断 `arr[i] == arr[i+n/4]` 是否相等即可。

## 代码

```go
func findSpecialInteger(arr []int) int {
	n := len(arr)
	for i := 0; i < n-n/4; i++ {
		if arr[i] == arr[i+n/4] {
			return arr[i]
		}
	}
	return -1
}
```