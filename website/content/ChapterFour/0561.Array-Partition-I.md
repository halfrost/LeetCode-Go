# [561. Array Partition I](https://leetcode.com/problems/array-partition-i/)


## 题目

Given an array of **2n** integers, your task is to group these integers into **n** pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

**Example 1**:

```
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
```

**Note**:

1. **n** is a positive integer, which is in the range of [1, 10000].
2. All the integers in the array will be in the range of [-10000, 10000].

## 题目大意

给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。


## 解题思路

- 给定一个 2n 个数组，要求把它们分为 n 组一行，求出各组最小值的总和的最大值。
- 由于题目给的数据范围不大，[-10000, 10000]，所以我们可以考虑用一个哈希表数组，里面存储 i - 10000 元素的频次，偏移量是 10000。这个哈希表能按递增的顺序访问数组，这样可以减少排序的耗时。题目要求求出分组以后求和的最大值，那么所有偏小的元素尽量都安排在一组里面，这样取 min 以后，对最大和影响不大。例如，(1 , 1) 这样安排在一起，min 以后就是 1 。但是如果把相差很大的两个元素安排到一起，那么较大的那个元素就“牺牲”了。例如，(1 , 10000)，取 min 以后就是 1，于是 10000 就“牺牲”了。所以需要优先考虑较小值。
- 较小值出现的频次可能是奇数也可能是偶数。如果是偶数，那比较简单，把它们俩俩安排在一起就可以了。如果是奇数，那么它会落单一次，落单的那个需要和距离它最近的一个元素进行配对，这样对最终的和影响最小。较小值如果是奇数，那么就会影响后面元素的选择，后面元素如果是偶数，由于需要一个元素和前面的较小值配对，所以它剩下的又是奇数个。这个影响会依次传递到后面。所以用一个 flag 标记，如果当前集合中有剩余元素将被再次考虑，则此标志设置为 1。在从下一组中选择元素时，会考虑已考虑的相同额外元素。
- 最后扫描过程中动态的维护 sum 值就可以了。

## 代码

```go

package leetcode

func arrayPairSum(nums []int) int {
	array := [20001]int{}
	for i := 0; i < len(nums); i++ {
		array[nums[i]+10000]++
	}
	flag, sum := true, 0
	for i := 0; i < len(array); i++ {
		for array[i] > 0 {
			if flag {
				sum = sum + i - 10000
			}
			flag = !flag
			array[i]--
		}
	}
	return sum
}

```