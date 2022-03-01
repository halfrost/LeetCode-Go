# [2167. Minimum Time to Remove All Cars Containing Illegal Goods](https://leetcode.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/)


## 题目

You are given a **0-indexed** binary string `s` which represents a sequence of train cars. `s[i] = '0'` denotes that the `ith` car does **not** contain illegal goods and `s[i] = '1'` denotes that the `ith` car does contain illegal goods.

As the train conductor, you would like to get rid of all the cars containing illegal goods. You can do any of the following three operations **any** number of times:

1. Remove a train car from the **left** end (i.e., remove `s[0]`) which takes 1 unit of time.
2. Remove a train car from the **right** end (i.e., remove `s[s.length - 1]`) which takes 1 unit of time.
3. Remove a train car from **anywhere** in the sequence which takes 2 units of time.

Return *the **minimum** time to remove all the cars containing illegal goods*.

Note that an empty sequence of cars is considered to have no cars containing illegal goods.

**Example 1:**

```
Input: s = "1100101"
Output: 5
Explanation:
One way to remove all the cars containing illegal goods from the sequence is to
- remove a car from the left end 2 times. Time taken is 2 * 1 = 2.
- remove a car from the right end. Time taken is 1.
- remove the car containing illegal goods found in the middle. Time taken is 2.
This obtains a total time of 2 + 1 + 2 = 5.

An alternative way is to
- remove a car from the left end 2 times. Time taken is 2 * 1 = 2.
- remove a car from the right end 3 times. Time taken is 3 * 1 = 3.
This also obtains a total time of 2 + 3 = 5.

5 is the minimum time taken to remove all the cars containing illegal goods.
There are no other ways to remove them with less time.

```

**Example 2:**

```
Input: s = "0010"
Output: 2
Explanation:
One way to remove all the cars containing illegal goods from the sequence is to
- remove a car from the left end 3 times. Time taken is 3 * 1 = 3.
This obtains a total time of 3.

Another way to remove all the cars containing illegal goods from the sequence is to
- remove the car containing illegal goods found in the middle. Time taken is 2.
This obtains a total time of 2.

Another way to remove all the cars containing illegal goods from the sequence is to
- remove a car from the right end 2 times. Time taken is 2 * 1 = 2.
This obtains a total time of 2.

2 is the minimum time taken to remove all the cars containing illegal goods.
There are no other ways to remove them with less time.
```

**Constraints:**

- `1 <= s.length <= 2 * 10^5`
- `s[i]` is either `'0'` or `'1'`.

## 题目大意

给你一个下标从 0 开始的二进制字符串 s ，表示一个列车车厢序列。s[i] = '0' 表示第 i 节车厢 不 含违禁货物，而 s[i] = '1' 表示第 i 节车厢含违禁货物。

作为列车长，你需要清理掉所有载有违禁货物的车厢。你可以不限次数执行下述三种操作中的任意一个：

1. 从列车 左 端移除一节车厢（即移除 s[0]），用去 1 单位时间。
2. 从列车 右 端移除一节车厢（即移除 s[s.length - 1]），用去 1 单位时间。
3. 从列车车厢序列的 任意位置 移除一节车厢，用去 2 单位时间。

返回移除所有载有违禁货物车厢所需要的 最少 单位时间数。注意，空的列车车厢序列视为没有车厢含违禁货物。

## 解题思路

- 这道题求最少单位时间数，最少时间数一定是尽量少使用 2 个单位时间的操作，多用 1 个时间的操作。从列车两头移除车厢，只需要移除和旁边车厢的金属连接处即可。由于列车位于两边，所以与其他车厢的金属连接处只有 1 个，故只需要 1 个单位时间；当车厢在中间，该车厢与两边的车厢有 2 个金属连接处，移除它需要断开与两边车厢的连接。所以需要 2 个单位时间。
- 断开中间一节车厢以后，列车会被断成 2 部分。2 部分列车分别有 2 个头 2 个尾。举例：`1100111101`，如果把它从第 5 节开始断开，剩下的列车为 `11001 (1)` 和 `1101`。剩下的 1 都位于 2 边，移除他们都只需要 1 个单位时间。那么移除所有违禁品最少时间是 2 * 1 + 1 * 6 = 8。
- 左半部分，定义 prefixSum[i] 表示移除前 i 节车厢所花费的最少时间。状态转移方程为：
    
    $prefixSum[i] =\left\{\begin{matrix}prefixSum[i-1],s[i]=0\\ min(prefixSum[i-1]+2, i+1), s[i]=1\end{matrix}\right.$
    
- 同理，右半部分定义 suffixSum[i] 表示移除后 i 节车厢所花费的最少时间。状态转移方程为：
    
    $suffixSum[i] =\left\{\begin{matrix} suffixSum[i+1],s[i]=0\\ min(suffixSum[i+1]+2, n-i), s[i]=1\end{matrix}\right.$
    
- 最后一层循环枚举 prefixSum[i] + suffixSum[i+1] 的最小值即为答案。
- 这一题在解法一的基础上还可以再简化。当 s[i] = 1 时，prefixSum 和 suffixSum 是两种计算方法。我们可以假设中间断开的部分在 prefixSum 中。于是可以合并上面两个状态转移方程。简化以后的代码见解法二。

## 代码

```go
package leetcode

import "runtime/debug"

// 解法一 DP
func minimumTime(s string) int {
	suffixSum, prefixSum, res := make([]int, len(s)+1), make([]int, len(s)+1), 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			suffixSum[i] = suffixSum[i+1]
		} else {
			suffixSum[i] = min(suffixSum[i+1]+2, len(s)-i)
		}
	}
	res = suffixSum[0]
	if s[0] == '1' {
		prefixSum[0] = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			prefixSum[i] = prefixSum[i-1]
		} else {
			prefixSum[i] = min(prefixSum[i-1]+2, i+1)
		}
		res = min(res, prefixSum[i]+suffixSum[i+1])
	}
	return res
}

func init() { debug.SetGCPercent(-1) }

// 解法二 小幅优化时间和空间复杂度
func minimumTime1(s string) int {
	res, count := len(s), 0
	for i := 0; i < len(s); i++ {
		count = min(count+int(s[i]-'0')*2, i+1)
		res = min(res, count+len(s)-i-1)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
```