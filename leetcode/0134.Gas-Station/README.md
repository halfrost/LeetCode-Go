# [134. Gas Stations](https://leetcode.com/problems/gas-station/)


## 题目

You are given a list of gas stations on a circular route, where the amount of gas at station i is gas[i], and the cost to travel from station i to its next station (i+1) is cost[i]. You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Write a function that returns the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.

**Example 1:**

```
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation: Start at station 3 (index 3), you have 4 liters of gas. Now you have a total of =0+4=4 liters of gas.
Travel to station 4. Now you have 4 - 1 + 5 = 8 liters of gas.
Travel to station 0. Now you have 8 - 2 + 1 = 7 liters of gas.
Travel to station 1. Now you have 7 - 3 + 2 = 6 liters of gas.
Travel to station 2. Now you have 6 - 4 + 3 = 5 liters of gas.
Travel to station 3 again. This time you will have consumed 5 liters of gas, which is exactly enough for you to return to station 3.
Therefore, 3 can be the starting index.
```

**Example 2:**

```
Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation: You cannot start at station 0 or 1 because there is not enough gas to get you to the next station.
We start at station 2, we get 4 liters of gas. Now you have a total of =0+4=4 liters of gas.
Travel to station 0. Now you have 4 - 3 + 2 = 3 liters of gas.
Travel to station 1. Now you have 3 - 3 + 3 = 3 liters of gas.
You cannot return to station 2 because it requires 4 liters of gas but you only have 3 liters in your tank.
Therefore, no matter what, you cannot travel around the ring.
```

**Constraints:**

- `gas.length == n`
- `cost.length == n`
- `1 <= n <= 10^5`
- `0 <= gas[i], cost[i] <= 10^4`

## 题目大意

在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

## 解题思路

- 初始化总油量（totalGas）和总消耗（totalCost）为0，起始点（start）为0，当前油量（currGas）为0。
  遍历每个加油站，累加油量和消耗，计算当前油量。
  如果当前油量小于0，说明从上一个加油站到当前加油站的油量不足以到达下一个加油站，因此将起始点设置为当前加油站的下一个位置，并将当前油量重置为0。
  遍历结束后，如果总油量小于总消耗，说明无法完成整个旅程，返回-1；否则返回起始点。

## 代码

```go
package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currGas += gas[i] - cost[i]

		if currGas < 0 {
			start = i + 1
			currGas = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

```