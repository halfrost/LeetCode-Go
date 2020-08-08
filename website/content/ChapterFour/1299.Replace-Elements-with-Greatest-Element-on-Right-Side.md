# [1299. Replace Elements with Greatest Element on Right Side](https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/)



## 题目

Given an array `arr`, replace every element in that array with the greatest element among the elements to its right, and replace the last element with `-1`.

After doing so, return the array.

**Example 1**:

```
Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]
```

**Constraints**:

- `1 <= arr.length <= 10^4`
- `1 <= arr[i] <= 10^5`


## 题目大意

给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。完成所有替换操作后，请你返回这个数组。

提示：

- 1 <= arr.length <= 10^4
- 1 <= arr[i] <= 10^5


## 解题思路

- 给出一个数组，要求把所有元素都替换成自己右边最大的元素，最后一位补上 -1 。最后输出变化以后的数组。
- 简单题，按照题意操作即可。

## 代码

```go
func replaceElements(arr []int) []int {
	j, temp := -1, 0
	for i := len(arr) - 1; i >= 0; i-- {
		temp = arr[i]
		arr[i] = j
		j = max(j, temp)
	}
	return arr
}
```