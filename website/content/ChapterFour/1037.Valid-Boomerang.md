# [1037. Valid Boomerang](https://leetcode.com/problems/valid-boomerang/)


## 题目

A *boomerang* is a set of 3 points that are all distinct and **not** in a straight line.

Given a list of three points in the plane, return whether these points are a boomerang.

**Example 1**:

```
Input: [[1,1],[2,3],[3,2]]
Output: true
```

**Example 2**:

```
Input: [[1,1],[2,2],[3,3]]
Output: false
```

**Note**:

1. `points.length == 3`
2. `points[i].length == 2`
3. `0 <= points[i][j] <= 100`

## 题目大意

回旋镖定义为一组三个点，这些点各不相同且不在一条直线上。给出平面上三个点组成的列表，判断这些点是否可以构成回旋镖。

## 解题思路

- 判断给出的 3 组点能否满足回旋镖。
- 简单题。判断 3 个点组成的 2 条直线的斜率是否相等。由于斜率的计算是除法，还可能遇到分母为 0 的情况，那么可以转换成乘法，交叉相乘再判断是否相等，就可以省去判断分母为 0 的情况了，代码也简洁成一行了。

## 代码

```go

package leetcode

func isBoomerang(points [][]int) bool {
	return (points[0][0]-points[1][0])*(points[0][1]-points[2][1]) != (points[0][0]-points[2][0])*(points[0][1]-points[1][1])
}

```