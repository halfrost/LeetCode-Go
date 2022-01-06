# [825. Friends Of Appropriate Ages](https://leetcode.com/problems/friends-of-appropriate-ages/)


## 题目

There are `n` persons on a social media website. You are given an integer array `ages` where `ages[i]` is the age of the `ith` person.

A Person `x` will not send a friend request to a person `y` (`x != y`) if any of the following conditions is true:

- `age[y] <= 0.5 * age+ 7`
- `age[y] > age[x]`
- `age[y] > 100 && age< 100`

Otherwise, `x` will send a friend request to `y`.

Note that if `x` sends a request to `y`, `y` will not necessarily send a request to `x`. Also, a person will not send a friend request to themself.

Return *the total number of friend requests made*.

**Example 1:**

```
Input: ages = [16,16]
Output: 2
Explanation: 2 people friend request each other.

```

**Example 2:**

```
Input: ages = [16,17,18]
Output: 2
Explanation: Friend requests are made 17 -> 16, 18 -> 17.

```

**Example 3:**

```
Input: ages = [20,30,100,110,120]
Output: 3
Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.

```

**Constraints:**

- `n == ages.length`
- `1 <= n <= 2 * 10^4`
- `1 <= ages[i] <= 120`

## 题目大意

在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。

如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：

- ages[y] <= 0.5 * ages[x] + 7
- ages[y] > ages[x]
- ages[y] > 100 && ages[x] < 100

否则，x 将会向 y 发送一条好友请求。注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。返回在该社交媒体网站上产生的好友请求总数。

## 解题思路

- 解法三，暴力解法。先统计 [1,120] 范围内每个年龄的人数。然后利用题目中的三个判断条件，筛选符合条件的用户对。需要注意的是，相同年龄的人可以相互发送好友请求。不同年龄的人发送好友请求是单向的，即年龄老的向年龄轻的发送好友请求，年龄轻的不会对年龄老的发送好友请求。
- 解法二，排序 + 双指针。题目给定的 3 个条件其实是 2 个。条件 3 包含在条件 2 中。条件 1 和条件 2 组合起来是 `0.5 × ages[x]+7 < ages[y] ≤ ages[x]`。当 ages[x] 小于 15 时，这个等式无解。考虑到年龄是单调递增的，`(0.5 × ages[x]+7,ages[x]]` 这个区间左右边界也是单调递增的。于是可以用双指针维护两个边界。在区间 [left, right] 内，这些下标对应的的 y 值都满足条件。当 `ages[left] > 0.5 × ages[x]+7` 时，左指针停止右移。当 `ages[right+1] > ages[x]` 时， 右指针停止右移。在 `[left, right]` 区间内，满足条件的 y 有 `right-left+1` 个，即使得 `ages[y]` 取值在 `(0.5 × ages[x]+7,ages[x]]` 之间。依照题意，`x≠y`，即该区间右边界取不到。y 的取值个数需要再减一，减去的是取到和 x 相同的值的下标。那么每个区间能取 `right-left` 个值。累加所有满足条件的值即为好友请求总数。
- 解法一。在解法二中，计算满足不等式 y 下标所在区间的时候，区间和区间存在重叠的情况，这些重叠情况导致了重复计算。所以这里可以优化。可以用 prefix sum 前缀和数组优化。代码见下方。

## 代码

```go
package leetcocde

import "sort"

// 解法一 前缀和，时间复杂度 O(n)
func numFriendRequests(ages []int) int {
	count, prefixSum, res := make([]int, 121), make([]int, 121), 0
	for _, age := range ages {
		count[age]++
	}
	for i := 1; i < 121; i++ {
		prefixSum[i] = prefixSum[i-1] + count[i]
	}
	for i := 15; i < 121; i++ {
		if count[i] > 0 {
			bound := i/2 + 8
			res += count[i] * (prefixSum[i] - prefixSum[bound-1] - 1)
		}
	}
	return res
}

// 解法二 双指针 + 排序，时间复杂度 O(n logn)
func numFriendRequests1(ages []int) int {
	sort.Ints(ages)
	left, right, res := 0, 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		res += right - left
	}
	return res
}

// 解法三 暴力解法 O(n^2)
func numFriendRequests2(ages []int) int {
	res, count := 0, [125]int{}
	for _, x := range ages {
		count[x]++
	}
	for i := 1; i <= 120; i++ {
		for j := 1; j <= 120; j++ {
			if j > i {
				continue
			}
			if (j-7)*2 <= i {
				continue
			}
			if j > 100 && i < 100 {
				continue
			}
			if i != j {
				res += count[i] * count[j]
			} else {
				res += count[i] * (count[j] - 1)
			}
		}
	}
	return res
}
```