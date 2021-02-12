# [119. Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)


## 题目

Given an integer `rowIndex`, return the `rowIndexth` row of the Pascal's triangle.

Notice that the row index starts from **0**.

![https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

In Pascal's triangle, each number is the sum of the two numbers directly above it.

**Follow up:**

Could you optimize your algorithm to use only *O*(*k*) extra space?

**Example 1:**

```
Input: rowIndex = 3
Output: [1,3,3,1]
```

**Example 2:**

```
Input: rowIndex = 0
Output: [1]
```

**Example 3:**

```
Input: rowIndex = 1
Output: [1,1]
```

**Constraints:**

- `0 <= rowIndex <= 33`

## 题目大意

给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

## 解题思路

- 题目中的三角是杨辉三角，每个数字是 `(a+b)^n` 二项式展开的系数。题目要求我们只能使用 O(k) 的空间。那么需要找到两两项直接的递推关系。由组合知识得知：

    $$\begin{aligned}C_{n}^{m} &= \frac{n!}{m!(n-m)!} \\C_{n}^{m-1} &= \frac{n!}{(m-1)!(n-m+1)!}\end{aligned}$$

    于是得到递推公式：

    $$C_{n}^{m} = C_{n}^{m-1} \times \frac{n-m+1}{m}$$

    利用这个递推公式即可以把空间复杂度优化到 O(k)

## 代码

```go
package leetcode

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
```