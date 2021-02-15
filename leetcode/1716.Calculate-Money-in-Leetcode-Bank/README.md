# [1716. Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/)


## 题目

Hercy wants to save money for his first car. He puts money in the Leetcode bank **every day**.

He starts by putting in `$1` on Monday, the first day. Every day from Tuesday to Sunday, he will put in `$1` more than the day before. On every subsequent Monday, he will put in `$1` more than the **previous Monday**.

Given `n`, return *the total amount of money he will have in the Leetcode bank at the end of the* `nth` *day.*

**Example 1:**

```
Input: n = 4
Output: 10
Explanation: After the 4th day, the total is 1 + 2 + 3 + 4 = 10.
```

**Example 2:**

```
Input: n = 10
Output: 37
Explanation: After the 10th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4) = 37. Notice that on the 2nd Monday, Hercy only puts in $2.
```

**Example 3:**

```
Input: n = 20
Output: 96
Explanation: After the 20th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4 + 5 + 6 + 7 + 8) + (3 + 4 + 5 + 6 + 7 + 8) = 96.
```

**Constraints:**

- `1 <= n <= 1000`

## 题目大意

Hercy 想要为购买第一辆车存钱。他 每天 都往力扣银行里存钱。最开始，他在周一的时候存入 1 块钱。从周二到周日，他每天都比前一天多存入 1 块钱。在接下来每一个周一，他都会比 前一个周一 多存入 1 块钱。给你 n ，请你返回在第 n 天结束的时候他在力扣银行总共存了多少块钱。

## 解题思路

- 简单题。按照题意写 2 层循环即可。

## 代码

```go
package leetcode

func totalMoney(n int) int {
	res := 0
	for tmp, count := 1, 7; n > 0; tmp, count = tmp+1, 7 {
		for m := tmp; n > 0 && count > 0; m++ {
			res += m
			n--
			count--
		}
	}
	return res
}
```