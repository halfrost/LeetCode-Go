# [1725. Number Of Rectangles That Can Form The Largest Square](https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/)


## 题目

You are given an array `rectangles` where `rectangles[i] = [li, wi]` represents the `ith` rectangle of length `li` and width `wi`.

You can cut the `ith` rectangle to form a square with a side length of `k` if both `k <= li` and `k <= wi`. For example, if you have a rectangle `[4,6]`, you can cut it to get a square with a side length of at most `4`.

Let `maxLen` be the side length of the **largest** square you can obtain from any of the given rectangles.

Return *the **number** of rectangles that can make a square with a side length of* `maxLen`.

**Example 1:**

```
Input: rectangles = [[5,8],[3,9],[5,12],[16,5]]
Output: 3
Explanation: The largest squares you can get from each rectangle are of lengths [5,3,5,5].
The largest possible square is of length 5, and you can get it out of 3 rectangles.
```

**Example 2:**

```
Input: rectangles = [[2,3],[3,7],[4,3],[3,7]]
Output: 3
```

**Constraints:**

- `1 <= rectangles.length <= 1000`
- `rectangles[i].length == 2`
- `1 <= li, wi <= 10^9`
- `li != wi`

## 题目大意

给你一个数组 rectangles ，其中 rectangles[i] = [li, wi] 表示第 i 个矩形的长度为 li 、宽度为 wi 。如果存在 k 同时满足 k <= li 和 k <= wi ，就可以将第 i 个矩形切成边长为 k 的正方形。例如，矩形 [4,6] 可以切成边长最大为 4 的正方形。设 maxLen 为可以从矩形数组 rectangles 切分得到的 最大正方形 的边长。返回可以切出边长为 maxLen 的正方形的矩形 数目 。

## 解题思路

- 简单题。扫描数组中的每一个矩形，先找到边长较小的那条边，作为正方形的边长。扫描过程中动态更新最大的正方形边长，并累加计数。循环一遍结束，输出最终计数值即可。

## 代码

```go
package leetcode

func countGoodRectangles(rectangles [][]int) int {
	minLength, count := 0, 0
	for i, _ := range rectangles {
		minSide := 0
		if rectangles[i][0] <= rectangles[i][1] {
			minSide = rectangles[i][0]
		} else {
			minSide = rectangles[i][1]
		}
		if minSide > minLength {
			minLength = minSide
			count = 1
		} else if minSide == minLength {
			count++
		}
	}
	return count
}
```