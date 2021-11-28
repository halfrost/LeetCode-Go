# [495. Teemo Attacking](https://leetcode.com/problems/teemo-attacking/)


## 题目

Our hero Teemo is attacking an enemy Ashe with poison attacks! When Teemo attacks Ashe, Ashe gets poisoned for a exactly duration seconds. 

More formally, an attack at second t will mean Ashe is poisoned during the inclusive time interval [t, t + duration - 1].

If Teemo attacks again before the poison effect ends, the timer for it is reset, and the poison effect will end duration seconds after the new attack.

You are given a non-decreasing integer array timeSeries, where timeSeries[i] denotes that Teemo attacks Ashe at second timeSeries[i], and an integer duration.

Return the total number of seconds that Ashe is poisoned.

**Example 1**:
```
Input: timeSeries = [1,4], duration = 2
Output: 4
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 4, Teemo attacks, and Ashe is poisoned for seconds 4 and 5.
Ashe is poisoned for seconds 1, 2, 4, and 5, which is 4 seconds in total.
```

**Example 2**:
```
Input: timeSeries = [1,2], duration = 2
Output: 3
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 2 however, Teemo attacks again and resets the poison timer. Ashe is poisoned for seconds 2 and 3.
Ashe is poisoned for seconds 1, 2, and 3, which is 3 seconds in total.
```

**Constraints**:

- 1 <= timeSeries.length <= 10000
- 0 <= timeSeries[i], duration <= 10000000
- timeSeries is sorted in non-decreasing order.

## 题目大意

在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄。他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。

当提莫攻击艾希，艾希的中毒状态正好持续duration 秒。

正式地讲，提莫在t发起发起攻击意味着艾希在时间区间 [t, t + duration - 1]（含 t 和 t + duration - 1）处于中毒状态。

如果提莫在中毒影响结束前再次攻击，中毒状态计时器将会重置，在新的攻击之后，中毒影响将会在duration秒后结束。

给你一个非递减的整数数组timeSeries，其中timeSeries[i]表示提莫在timeSeries[i]秒时对艾希发起攻击，以及一个表示中毒持续时间的整数duration 。

返回艾希处于中毒状态的总秒数。

## 解题思路

- i 从 1 开始计数，令 t 等于 timeSeries[i - 1]
- 比较 end(t + duration - 1) 和 timeSeries[i] 的大小，
  - 如果 end 小于 timeSeries[i],ans+=duration
  - 否则 ans += timeSeries[i] - t
- ans += duration 并返回 ans

## 代码

```go

package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	var ans int
	for i := 1; i < len(timeSeries); i++ {
		t := timeSeries[i-1]
		end := t + duration - 1
		if end < timeSeries[i] {
			ans += duration
		} else {
			ans += timeSeries[i] - t
		}
	}
	ans += duration
	return ans
}

```
