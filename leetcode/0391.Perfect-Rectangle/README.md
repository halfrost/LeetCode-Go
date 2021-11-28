# [391. Perfect Rectangle](https://leetcode.com/problems/perfect-rectangle/)

## 题目

Given an array rectangles where rectangles[i] = [xi, yi, ai, bi] represents an axis-aligned rectangle. The bottom-left point of the rectangle is (xi, yi) and the top-right point of it is (ai, bi).

Return true if all the rectangles together form an exact cover of a rectangular region.

**Example1:**

![https://assets.leetcode.com/uploads/2021/03/27/perectrec1-plane.jpg](https://assets.leetcode.com/uploads/2021/03/27/perectrec1-plane.jpg)

    Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
    Output: true
    Explanation: All 5 rectangles together form an exact cover of a rectangular region.

**Example2:**

![https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg](https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg)

    Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
    Output: false
    Explanation: Because there is a gap between the two rectangular regions.

**Example3:**

![https://assets.leetcode.com/uploads/2021/03/27/perfectrec3-plane.jpg](https://assets.leetcode.com/uploads/2021/03/27/perfectrec3-plane.jpg)

    Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]
    Output: false
    Explanation: Because there is a gap in the top center.

**Example4:**

![https://assets.leetcode.com/uploads/2021/03/27/perfecrrec4-plane.jpg](https://assets.leetcode.com/uploads/2021/03/27/perfecrrec4-plane.jpg)

    Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
    Output: false
    Explanation: Because two of the rectangles overlap with each other.

**Constraints:**

- 1 <= rectangles.length <= 2 * 10000
- rectangles[i].length == 4
- -100000 <= xi, yi, ai, bi <= 100000

## 题目大意

给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。

如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。

## 解题思路

- 矩形区域的面积等于所有矩形的面积之和并且满足矩形区域四角的顶点只能出现一次，且其余顶点的出现次数只能是两次或四次则返回 true,否则返回 false

## 代码

```go

package leetcode

type point struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {
	minX, minY, maxA, maxB := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	area := 0
	cnt := make(map[point]int)
	for _, v := range rectangles {
		x, y, a, b := v[0], v[1], v[2], v[3]
		area += (a - x) * (b - y)
		minX = min(minX, x)
		minY = min(minY, y)
		maxA = max(maxA, a)
		maxB = max(maxB, b)
		cnt[point{x, y}]++
		cnt[point{a, b}]++
		cnt[point{x, b}]++
		cnt[point{a, y}]++
	}
	if area != (maxA - minX) * (maxB - minY) ||
		cnt[point{minX, minY}] != 1 || cnt[point{maxA, maxB}] != 1 ||
		cnt[point{minX, maxB}] != 1 || cnt[point{maxA, minY}] != 1 {
		return false
	}
	delete(cnt, point{minX, minY})
	delete(cnt, point{maxA, maxB})
	delete(cnt, point{minX, maxB})
	delete(cnt, point{maxA, minY})
	for _, v := range cnt {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```