# [2021. Brightest Position on Street](https://leetcode.com/problems/brightest-position-on-street/)


## 题目

A perfectly straight street is represented by a number line. The street has street lamp(s) on it and is represented by a 2D integer array `lights`. Each `lights[i] = [positioni, rangei]` indicates that there is a street lamp at position `positioni` that lights up the area from `[positioni - rangei, positioni + rangei]` (**inclusive**).

The **brightness** of a position `p` is defined as the number of street lamp that light up the position `p`.

Given `lights`, return *the **brightest** position on the street. If there are multiple brightest positions, return the **smallest** one.*

**Example 1:**

![https://assets.leetcode.com/uploads/2021/09/28/image-20210928155140-1.png](https://assets.leetcode.com/uploads/2021/09/28/image-20210928155140-1.png)

```
Input: lights = [[-3,2],[1,2],[3,3]]
Output: -1
Explanation:
The first street lamp lights up the area from [(-3) - 2, (-3) + 2] = [-5, -1].
The second street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].
The third street lamp lights up the area from [3 - 3, 3 + 3] = [0, 6].

Position -1 has a brightness of 2, illuminated by the first and second street light.
Positions 0, 1, 2, and 3 have a brightness of 2, illuminated by the second and third street light.
Out of all these positions, -1 is the smallest, so return it.

```

**Example 2:**

```
Input: lights = [[1,0],[0,1]]
Output: 1
Explanation:
The first street lamp lights up the area from [1 - 0, 1 + 0] = [1, 1].
The second street lamp lights up the area from [0 - 1, 0 + 1] = [-1, 1].

Position 1 has a brightness of 2, illuminated by the first and second street light.
Return 1 because it is the brightest position on the street.

```

**Example 3:**

```
Input: lights = [[1,2]]
Output: -1
Explanation:
The first street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].

Positions -1, 0, 1, 2, and 3 have a brightness of 1, illuminated by the first street light.
Out of all these positions, -1 is the smallest, so return it.

```

**Constraints:**

- `1 <= lights.length <= 105`
- `lights[i].length == 2`
- `108 <= positioni <= 108`
- `0 <= rangei <= 108`

## 题目大意

一条完全笔直的街道由一条数字线表示。街道上有路灯，由二维数据表示。每个 `lights[i] = [positioni, rangei]` 表示位置 `i` 处有一盏路灯，灯可以照亮从 `[positioni - rangei, positioni + rangei]` （含）的区域。 位置 `p` 的亮度定义为点亮位置 `p` 的路灯数量。 给定路灯，返回街道上最亮的位置。如果有多个最亮的位置，则返回最小的一个。

## 解题思路

- 先将每个路灯的起始和终点位置计算出来。这样我们得到了一堆坐标点。假设灯照亮的范围是 [A, B]，那么在坐标轴上 A 坐标点处 + 1， B + 1 坐标点处 -1 。这样处理的含义是：坐标点 A 可以被一盏灯照亮，所以它照亮次数加一，坐标点 B + 1 出了灯照亮的范围了，所以照亮次数减一。那么从坐标轴坐标开始扫一遍，每次遇到 + 1 的时候就 + 1，遇到 - 1 的地方就 - 1。如此可以算出某个坐标点处，可以被灯照亮的总次数。
- 需要注意的点是，题目给的测试数据可能会有单点照亮的情况，即某一盏灯只照亮一个坐标点，灯照范围为 0。同一个坐标点也可能是多个灯的起点。用一个 map 去重坐标点即可。

## 代码

```go
package leetcode

import (
	"sort"
)

type lightItem struct {
	index int
	sign  int
}

func brightestPosition(lights [][]int) int {
	lightMap, lightItems := map[int]int{}, []lightItem{}
	for _, light := range lights {
		lightMap[light[0]-light[1]] += 1
		lightMap[light[0]+light[1]+1] -= 1
	}
	for k, v := range lightMap {
		lightItems = append(lightItems, lightItem{index: k, sign: v})
	}
	sort.SliceStable(lightItems, func(i, j int) bool {
		return lightItems[i].index < lightItems[j].index
	})
	res, border, tmp := 0, 0, 0
	for _, v := range lightItems {
		tmp += v.sign
		if border < tmp {
			res = v.index
			border = tmp
		}
	}
	return res
}
```