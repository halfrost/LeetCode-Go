# [1296. Divide Array in Sets of K Consecutive Numbers](https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/)

## 题目

Given an array of integers nums and a positive integer k, check whether it is possible to divide this array into sets of k consecutive numbers.

Return true if it is possible. Otherwise, return false.

**Example 1**:

    Input: nums = [1,2,3,3,4,4,5,6], k = 4
    Output: true
    Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].

**Example 2**:

    Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
    Output: true
    Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].

**Example 3**:

    Input: nums = [1,2,3,4], k = 3
    Output: false
    Explanation: Each array should be divided in subarrays of size 3.

**Constraints:**

- 1 <= k <= nums.length <= 100000
- 1 <= nums[i] <= 1000000000

## 题目大意

给你一个整数数组 nums 和一个正整数 k，请你判断是否可以把这个数组划分成一些由 k 个连续数字组成的集合。
如果可以，请返回 true；否则，返回 false。

## 解题思路

贪心算法

- 对nums升序排序
- 对nums内数字进行哈希计数（key:数字，value:数量）
- 遍历nums中的数字，以数量大于1的数字作为连续数字开头，寻找连续数字后续元素，若无法找到 k 个连续数字则返回false
- 所有数字都能找到 k 个连续数字返回true

##代码

```go
package leetcode

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] += 1
	}
	sort.Ints(nums)
	for _, num := range nums {
		if mp[num] == 0 {
			continue
		}
		for diff := 0; diff < k; diff++ {
			if mp[num+diff] == 0 {
				return false
			}
			mp[num+diff] -= 1
		}
	}
	return true
}
```
