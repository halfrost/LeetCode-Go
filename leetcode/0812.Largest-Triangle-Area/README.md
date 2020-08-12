# [812. Largest Triangle Area](https://leetcode.com/problems/largest-triangle-area/)


## 题目

You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

```
Example:
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2
Explanation: 
The five points are show in the figure below. The red triangle is the largest.
```

![https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/04/1027.png](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/04/1027.png)

**Notes**:

- `3 <= points.length <= 50`.
- No points will be duplicated.
- `-50 <= points[i][j] <= 50`.
- Answers within `10^-6` of the true value will be accepted as correct.

## 题目大意

给定包含多个点的集合，从其中取三个点组成三角形，返回能组成的最大三角形的面积。

## 解题思路

- 给出一组点的坐标，要求找出能组成三角形面积最大的点集合，输出这个最大面积。
- 数学题。按照数学定义，分别计算这些能构成三角形的点形成的三角形面积，最终输出最大面积即可。

## 代码

```go

package leetcode

func largestTriangleArea(points [][]int) float64 {
	maxArea, n := 0.0, len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				maxArea = max(maxArea, area(points[i], points[j], points[k]))
			}
		}
	}
	return maxArea
}

func area(p1, p2, p3 []int) float64 {
	return abs(p1[0]*p2[1]+p2[0]*p3[1]+p3[0]*p1[1]-p1[0]*p3[1]-p2[0]*p1[1]-p3[0]*p2[1]) / 2
}

func abs(num int) float64 {
	if num < 0 {
		num = -num
	}
	return float64(num)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

```