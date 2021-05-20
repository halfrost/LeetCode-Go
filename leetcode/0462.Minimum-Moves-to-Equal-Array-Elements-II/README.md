# [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)


## 题目

Given an integer array `nums` of size `n`, return *the minimum number of moves required to make all array elements equal*.

In one move, you can increment or decrement an element of the array by `1`.

**Example 1:**

```
Input: nums = [1,2,3]
Output: 2
Explanation:
Only two moves are needed (remember each move increments or decrements one element):
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
```

**Example 2:**

```
Input: nums = [1,10,2,9]
Output: 16
```

**Constraints:**

- `n == nums.length`
- `1 <= nums.length <= 10^5`
- `109 <= nums[i] <= 10^9`

## 题目大意

给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加 1 或减 1。 您可以假设数组的长度最多为10000。

## 解题思路

- 这题抽象成数学问题是，如果我们把数组 a 中的每个数看成水平轴上的一个点，那么根据上面的移动次数公式，我们需要找到在水平轴上找到一个点 x，使得这 N 个点到 x 的距离之和最小。有 2 个点值得我们考虑，一个是中位数，另外一个是平均值。举个简单的例子，[1,0,0,8,6] 这组数据，中位数是 1，平均值是 3 。分别计算移动的步数，按照中位数对齐是 14，按照平均值对齐是 16 。所以选择中位数。
- 此题可以用数学证明，证明出，按照平均值移动的步数 ≥ 按照中位数移动的步数。具体证明笔者这里不证明了，感兴趣的同学可以自己证明试试。

## 代码

```go
package leetcode

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	moves, mid := 0, len(nums)/2
	sort.Ints(nums)
	for i := range nums {
		if i == mid {
			continue
		}
		moves += int(math.Abs(float64(nums[mid] - nums[i])))
	}
	return moves
}
```