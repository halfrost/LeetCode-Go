# [1680. Concatenation of Consecutive Binary Numbers](https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/)


## 题目

Given an integer `n`, return *the **decimal value** of the binary string formed by concatenating the binary representations of* `1` *to* `n` *in order, **modulo*** `109 + 7`.

**Example 1:**

```
Input: n = 1
Output: 1
Explanation: "1" in binary corresponds to the decimal value 1. 
```

**Example 2:**

```
Input: n = 3
Output: 27
Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
After concatenating them, we have "11011", which corresponds to the decimal value 27.
```

**Example 3:**

```
Input: n = 12
Output: 505379714
Explanation: The concatenation results in "1101110010111011110001001101010111100".
The decimal value of that is 118505380540.
After modulo 109 + 7, the result is 505379714.
```

**Constraints:**

- `1 <= n <= 10^5`

## 题目大意

给你一个整数 n ，请你将 1 到 n 的二进制表示连接起来，并返回连接结果对应的 十进制 数字对 10^9 + 7 取余的结果。

## 解题思路

- 理解题意以后，先找到如何拼接最终二进制数的规律。假设 `f(n)` 为最终变换以后的十进制数。那么根据题意，`f(n) = f(n-1) << shift + n` 这是一个递推公式。`shift` 左移的位数就是 `n` 的二进制对应的长度。`shift` 的值是随着 `n` 变化而变化的。由二进制进位规律可以知道，2 的整数次幂的时候，对应的二进制长度会增加 1 位。这里可以利用位运算来判断是否是 2 的整数次幂。
- 这道题另外一个需要处理的是模运算的法则。此题需要用到模运算的加法法则。

    ```go
    模运算与基本四则运算有些相似，但是除法例外。
    (a + b) % p = (a % p + b % p) % p （1）
    (a - b) % p = (a % p - b % p) % p （2）
    (a * b) % p = (a % p * b % p) % p （3）
    a ^ b % p = ((a % p)^b) % p （4）
    结合律：
    ((a+b) % p + c) % p = (a + (b+c) % p) % p （5）
    ((a*b) % p * c)% p = (a * (b*c) % p) % p （6）
    交换律：
    (a + b) % p = (b+a) % p （7）
    (a * b) % p = (b * a) % p （8）
    分配律：
    ((a +b)% p * c) % p = ((a * c) % p + (b * c) % p) % p （9）
    ```

    这一题需要用到模运算的加法运算法则。

## 代码

```go
package leetcode

import (
	"math/bits"
)

// 解法一 模拟
func concatenatedBinary(n int) int {
	res, mod, shift := 0, 1000000007, 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			shift++
		}
		res = ((res << shift) + i) % mod
	}
	return res
}

// 解法二 位运算
func concatenatedBinary1(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res = (res<<bits.Len(uint(i)) | i) % (1e9 + 7)
	}
	return res
}
```