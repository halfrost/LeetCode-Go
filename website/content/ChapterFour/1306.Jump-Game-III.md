# [1306. Jump Game III](https://leetcode.com/problems/jump-game-iii/)


## 题目

Given an array of non-negative integers `arr`, you are initially positioned at `start` index of the array. When you are at index `i`, you can jump to `i + arr[i]` or `i - arr[i]`, check if you can reach to **any** index with value 0.

Notice that you can not jump outside of the array at any time.

**Example 1**:

```
Input: arr = [4,2,3,0,3,1,2], start = 5
Output: true
Explanation: 
All possible ways to reach at index 3 with value 0 are: 
index 5 -> index 4 -> index 1 -> index 3 
index 5 -> index 6 -> index 4 -> index 1 -> index 3
```

**Example 2**:

```
Input: arr = [4,2,3,0,3,1,2], start = 0
Output: true 
Explanation: 
One possible way to reach at index 3 with value 0 is: 
index 0 -> index 4 -> index 1 -> index 3
```

**Example 3**:

```
Input: arr = [3,0,2,1,2], start = 2
Output: false
Explanation: There is no way to reach at index 1 with value 0.
```

**Constraints**:

- `1 <= arr.length <= 5 * 10^4`
- `0 <= arr[i] < arr.length`
- `0 <= start < arr.length`


## 题目大意

这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。注意，不管是什么情况下，你都无法跳到数组之外。

提示：

- 1 <= arr.length <= 5 * 10^4
- 0 <= arr[i] < arr.length
- 0 <= start < arr.length


## 解题思路

- 给出一个非负数组和一个起始下标 `start`。站在 `start`，每次可以跳到 `i + arr[i]` 或者 `i - arr[i]` 。要求判断能否跳到元素值为 0 的下标处。
- 这一题考察的是递归。每一步都需要判断 3 种可能：
    1. 当前是否站在元素值为 0 的目标点上。
    2. 往前跳 arr[start]，是否能站在元素值为 0 的目标点上。
    3. 往后跳 arr[start]，是否能站在元素值为 0 的目标点上。

    第 2 种可能和第 3 种可能递归即可，每一步都判断这 3 种可能是否有一种能跳到元素值为 0 的下标处。

- `arr[start] += len(arr)`  这一步仅仅只是为了标记此下标已经用过了，下次判断的时候该下标会被 `if` 条件筛选掉。

## 代码

```go
func canReach(arr []int, start int) bool {
	if start >= 0 && start < len(arr) && arr[start] < len(arr) {
		jump := arr[start]
		arr[start] += len(arr)
		return jump == 0 || canReach(arr, start+jump) || canReach(arr, start-jump)
	}
	return false
}
```