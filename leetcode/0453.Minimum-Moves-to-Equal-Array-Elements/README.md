# [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)


## 题目

Given a **non-empty** integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

**Example**:

```
Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
```

## 题目大意

给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动将会使 n - 1 个元素增加 1。

## 解题思路

- 给定一个数组，要求输出让所有元素都相等的最小步数。每移动一步都会使得 n - 1 个元素 + 1 。
- 数学题。这道题正着思考会考虑到排序或者暴力的方法上去。反过来思考一下，使得每个元素都相同，意思让所有元素的差异变为 0 。每次移动的过程中，都有 n - 1 个元素 + 1，那么没有 + 1 的那个元素和其他 n - 1 个元素相对差异就缩小了。所以这道题让所有元素都变为相等的最少步数，即等于让所有元素相对差异减少到最小的那个数。想到这里，此题就可以优雅的解出来了。

## 代码

```go

package leetcode

import "math"

func minMoves(nums []int) int {
	sum, min, l := 0, math.MaxInt32, len(nums)
	for _, v := range nums {
		sum += v
		if min > v {
			min = v
		}
	}
	return sum - min*l
}

```