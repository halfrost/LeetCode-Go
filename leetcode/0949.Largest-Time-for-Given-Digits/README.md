# [949. Largest Time for Given Digits](https://leetcode.com/problems/largest-time-for-given-digits/)


## 题目

Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59. Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5. If no valid time can be made, return an empty string.

**Example 1**:

```
Input: [1,2,3,4]
Output: "23:41"
```

**Example 2**:

```
Input: [5,5,5,5]
Output: ""
```

**Note**:

1. `A.length == 4`
2. `0 <= A[i] <= 9`

## 题目大意

给定一个由 4 位数字组成的数组，返回可以设置的符合 24 小时制的最大时间。最小的 24 小时制时间是 00:00，而最大的是 23:59。从 00:00 （午夜）开始算起，过得越久，时间越大。以长度为 5 的字符串返回答案。如果不能确定有效时间，则返回空字符串。

## 解题思路

- 给出 4 个数字，要求返回一个字符串，代表由这 4 个数字能组成的最大 24 小时制的时间。
- 简单题，这一题直接暴力枚举就可以了。依次检查给出的 4 个数字每个排列组合是否是时间合法的。例如检查 10 * A[i] + A[j] 是不是小于 24， 10 * A[k] + A[l] 是不是小于 60。如果合法且比目前存在的最大时间更大，就更新这个最大时间。

## 代码

```go

package leetcode

import "fmt"

func largestTimeFromDigits(A []int) string {
	flag, res := false, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				l := 6 - i - j - k
				hour := A[i]*10 + A[j]
				min := A[k]*10 + A[l]
				if hour < 24 && min < 60 {
					if hour*60+min >= res {
						res = hour*60 + min
						flag = true
					}
				}
			}
		}
	}
	if flag {
		return fmt.Sprintf("%02d:%02d", res/60, res%60)
	} else {
		return ""
	}
}

```