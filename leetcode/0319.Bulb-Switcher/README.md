# [319. Bulb Switcher](https://leetcode.com/problems/bulb-switcher/)


## 题目

There are n bulbs that are initially off. You first turn on all the bulbs, then you turn off every second bulb.

On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on). For the ith round, you toggle every i bulb. For the nth round, you only toggle the last bulb.

Return the number of bulbs that are on after n rounds.

**Example 1:**

    Input: n = 3
    Output: 1
    Explanation: At first, the three bulbs are [off, off, off].
    After the first round, the three bulbs are [on, on, on].
    After the second round, the three bulbs are [on, off, on].
    After the third round, the three bulbs are [on, off, off].
    So you should return 1 because there is only one bulb is on.

**Example 2:**

    Input: n = 0
    Output: 0

**Example 3:**

    Input: n = 1
    Output: 1

## 题目大意

初始时有 n 个灯泡处于关闭状态。第一轮，你将会打开所有灯泡。接下来的第二轮，你将会每两个灯泡关闭一个。

第三轮，你每三个灯泡就切换一个灯泡的开关（即，打开变关闭，关闭变打开）。第 i 轮，你每 i 个灯泡就切换一个灯泡的开关。直到第 n 轮，你只需要切换最后一个灯泡的开关。

找出并返回 n 轮后有多少个亮着的灯泡。

## 解题思路

- 计算 1 到 n 中有奇数个约数的个数
- 1 到 n 中的某个数 x 有奇数个约数,也即 x 是完全平方数
- 计算 1 到 n 中完全平方数的个数 sqrt(n)

## 代码

```go

package leetcode

import "math"

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

```