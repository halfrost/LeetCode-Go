# [279. Perfect Squares](https://leetcode.com/problems/perfect-squares/)


## 题目

Given an integer `n`, return *the least number of perfect square numbers that sum to* `n`.

A **perfect square** is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, `1`, `4`, `9`, and `16` are perfect squares while `3` and `11` are not.

**Example 1:**

```
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
```

**Example 2:**

```
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
```

**Constraints:**

- `1 <= n <= 104`

## 题目大意

给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

## 解题思路

- 由拉格朗日的四平方定理可得，每个自然数都可以表示为四个整数平方之和。 其中四个数字是整数。四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和。这给出了本题的答案的上界。
- 四平方和定理可以推出三平方和推论：当且仅当 {{< katex >}} n \neq 4^{k} \times (8*m + 7){{< /katex >}} 时，n 可以被表示为至多三个正整数的平方和。所以当 {{< katex >}} n = 4^{k} * (8*m + 7){{< /katex >}}  时，n 只能被表示为四个正整数的平方和。此时我们可以直接返回 4。
- 当 {{< katex >}} n \neq 4^{k} \times (8*m + 7){{< /katex >}} 时，需要判断 n 到底可以分解成几个完全平方数之和。答案肯定是 1，2，3 中的一个。题目要求我们求最小的，所以从 1 开始一个个判断是否满足。如果答案为 1，代表 n 为完全平方数，这很好判断。如果答案为 2，代表 {{< katex >}} n = a^{2} + b^{2} {{< /katex >}}，枚举 {{< katex >}} 1 \leqslant a \leqslant \sqrt{n} {{< /katex >}}，判断  {{< katex >}} n - a^{2} {{< /katex >}} 是否为完全平方数。当 1 和 2 都排除了，剩下的答案只能为 3 了。

## 代码

```go
package leetcode

import "math"

func numSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

// 判断是否为完全平方数
func isPerfectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}
```