# [1710. Maximum Units on a Truck](https://leetcode.com/problems/maximum-units-on-a-truck/)


## 题目

You are assigned to put some amount of boxes onto **one truck**. You are given a 2D array `boxTypes`, where `boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]`:

- `numberOfBoxesi` is the number of boxes of type `i`.
- `numberOfUnitsPerBoxi`is the number of units in each box of the type `i`.

You are also given an integer `truckSize`, which is the **maximum** number of **boxes** that can be put on the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed `truckSize`.

Return *the **maximum** total number of **units** that can be put on the truck.*

**Example 1:**

```
Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
Output: 8
Explanation: There are:
- 1 box of the first type that contains 3 units.
- 2 boxes of the second type that contain 2 units each.
- 3 boxes of the third type that contain 1 unit each.
You can take all the boxes of the first and second types, and one box of the third type.
The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.
```

**Example 2:**

```
Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
Output: 91
```

**Constraints:**

- `1 <= boxTypes.length <= 1000`
- `1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000`
- `1 <= truckSize <= 106`

## 题目大意

请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：

- numberOfBoxesi 是类型 i 的箱子的数量。-
- numberOfUnitsPerBoxi 是类型 i 每个箱子可以装载的单元数量。

整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。返回卡车可以装载 单元 的 最大 总数。

## 解题思路

- 简单题。先将箱子按照单元数量从大到小排序。要想卡车装载单元数最大，那么需要尽量装单元数多的箱子。所以排序以后从单元数量多的箱子开始取。一直取至 truckSize 没有空间。累积的单元数即最大总数。

## 代码

```go
package leetcode

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for i := 0; truckSize > 0 && i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			truckSize -= boxTypes[i][0]
			res += (boxTypes[i][1] * boxTypes[i][0])
		} else {
			res += (truckSize * boxTypes[i][1])
			truckSize = 0
		}
	}
	return res
}
```