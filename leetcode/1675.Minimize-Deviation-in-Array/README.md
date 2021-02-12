# [1675. Minimize Deviation in Array](https://leetcode.com/problems/minimize-deviation-in-array/)

## 题目

You are given an array `nums` of `n` positive integers.

You can perform two types of operations on any element of the array any number of times:

- If the element is **even**, **divide** it by `2`.
    - For example, if the array is `[1,2,3,4]`, then you can do this operation on the last element, and the array will be `[1,2,3,2].`
- If the element is **odd**, **multiply** it by `2`.
    - For example, if the array is `[1,2,3,4]`, then you can do this operation on the first element, and the array will be `[2,2,3,4].`

The **deviation** of the array is the **maximum difference** between any two elements in the array.

Return *the **minimum deviation** the array can have after performing some number of operations.*

**Example 1:**

```
Input: nums = [1,2,3,4]
Output: 1
Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
```

**Example 2:**

```
Input: nums = [4,1,5,20,3]
Output: 3
Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
```

**Example 3:**

```
Input: nums = [2,10,8]
Output: 3
```

**Constraints:**

- `n == nums.length`
- `2 <= n <= 105`
- `1 <= nums[i] <= 10^9`

## 题目大意

给你一个由 n 个正整数组成的数组 nums 。你可以对数组的任意元素执行任意次数的两类操作：

- 如果元素是 偶数 ，除以 2。例如，如果数组是 [1,2,3,4] ，那么你可以对最后一个元素执行此操作，使其变成 [1,2,3,2]
- 如果元素是 奇数 ，乘上 2。例如，如果数组是 [1,2,3,4] ，那么你可以对第一个元素执行此操作，使其变成 [2,2,3,4]
数组的 偏移量 是数组中任意两个元素之间的 最大差值 。

返回数组在执行某些操作之后可以拥有的 最小偏移量 。

## 解题思路

- 要找到最小偏移量，即需要令最大值变小，最小值变大。要想达到这个要求，需要对奇数偶数做乘法和除法。这里特殊的是，奇数一旦乘以 2 以后，就变成偶数了。偶数除以 2 以后可能是奇数也可能是偶数。所以可以先将所有的奇数都乘以 2 统一变成偶数。
- 第二轮不断的将最大值除 2，直到最大值为奇数，不能再操作了。每轮循环中把比 min 值大的偶数也都除以 2 。这里除以 2 有 2 个目的，一个目的是将第一步奇数乘 2 还原回去，另一个目的是将本来的偶数除以 2 。可能有人有疑问，为什么只把最大值变小，没有将最小值变大呢？如果最小值是奇数，那么它一定是由上一个偶数除以 2 变过来的，我们在上一个状态已经计算过这个偶数了，因此没必要扩大它；如果最小值是偶数，那么它一定会在某一轮的除 2 操作中，不操作，即它不会满足 `min <= nums[i]/2`  这个条件。每次循环都更新该次循环的最大值和最小值，并记录偏移量。不断的循环，直到最大值为奇数，退出循环。最终输出最小偏移量。

## 代码

```go
package leetcode

func minimumDeviation(nums []int) int {
	min, max := 0, 0
	for i := range nums {
		if nums[i]%2 == 1 {
			nums[i] *= 2
		}
		if i == 0 {
			min = nums[i]
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		} else if max < nums[i] {
			max = nums[i]
		}
	}
	res := max - min
	for max%2 == 0 {
		tmax, tmin := 0, 0
		for i := range nums {
			if nums[i] == max || (nums[i]%2 == 0 && min <= nums[i]/2) {
				nums[i] /= 2
			}
			if i == 0 {
				tmin = nums[i]
				tmax = nums[i]
			} else if nums[i] < tmin {
				tmin = nums[i]
			} else if tmax < nums[i] {
				tmax = nums[i]
			}
		}
		if tmax-tmin < res {
			res = tmax - tmin
		}
		min, max = tmin, tmax
	}
	return res
}
```