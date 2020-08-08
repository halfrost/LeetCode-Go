# [1300. Sum of Mutated Array Closest to Target](https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/)



## 题目

Given an integer array `arr` and a target value `target`, return the integer `value` such that when we change all the integers larger than `value` in the given array to be equal to `value`, the sum of the array gets as close as possible (in absolute difference) to `target`.

In case of a tie, return the minimum such integer.

Notice that the answer is not neccesarilly a number from `arr`.

**Example 1**:

```
Input: arr = [4,9,3], target = 10
Output: 3
Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.
```

**Example 2**:

```
Input: arr = [2,3,5], target = 10
Output: 5
```

**Example 3**:

```
Input: arr = [60864,25176,27249,21296,20204], target = 56803
Output: 11361
```

**Constraints**:

- `1 <= arr.length <= 10^4`
- `1 <= arr[i], target <= 10^5`


## 题目大意

给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。请注意，答案不一定是 arr 中的数字。

提示：

- 1 <= arr.length <= 10^4
- 1 <= arr[i], target <= 10^5



## 解题思路

- 给出一个数组 arr 和 target。能否找到一个 value 值，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近 target。如果有多种方法，输出 value 值最小的。
- 这一题可以用二分搜索来求解。最后输出的唯一解有 2 个限制条件，一个是变化后的数组和最接近 target 。另一个是输出的 value 是所有可能方法中最小值。二分搜索最终的 value 值。mid 就是尝试的 value 值，每选择一次 mid，就算一次总和，和 target 比较。由于数组里面每个数和 mid 差距各不相同，所以每次调整 mid 有可能出现 mid 选小了以后，距离 target 反而更大了；mid 选大了以后，距离 target 反而更小了。这里的解决办法是，把 value 上下方可能的值都拿出来比较一下。

## 代码

```go
func findBestValue(arr []int, target int) int {
	low, high := 0, 100000
	for low < high {
		mid := low + (high-low)>>1
		if calculateSum(arr, mid) < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	// 比较阈值线分别定在 left - 1 和 left 的时候与 target 的接近程度
	sum1, sum2 := calculateSum(arr, low-1), calculateSum(arr, low)
	if target-sum1 <= sum2-target {
		return low - 1
	}
	return low
}

func calculateSum(arr []int, mid int) int {
	sum := 0
	for _, num := range arr {
		sum += min(num, mid)
	}
	return sum
}
```